# Feeding Mechanism Quick Reference

> **Last updated**: 2026-04-14 â€” added `auto_accept` and `promote_all_accepted`

---

## TL;DR â€” Fast Path (3 steps, any number of observations)

Start a **writable** MCP container on port 8092:

```powershell
$env:MCP_SERVER_PORT=8092; docker compose up -d --build mcp-server
```

**Step 1 â€” Ingest + auto-accept in one call**

```powershell
$r = Invoke-RestMethod -Uri "http://localhost:8092/mcp" -Method Post `
  -ContentType "application/json" -Body '{"jsonrpc":"2.0","id":1,"method":"feeding/ingest_batch","params":{"persist":true,"auto_accept":true,"continue_on_error":true}}'
"Accepted: $($r.result.accepted_count)"
```

**Step 2 â€” Promote all accepted at once**

```powershell
$r = Invoke-RestMethod -Uri "http://localhost:8092/mcp" -Method Post `
  -ContentType "application/json" -Body '{"jsonrpc":"2.0","id":2,"method":"feeding/promote_all_accepted","params":{}}'
"Promoted: $($r.result.promoted_count)"
```

**Step 3 â€” Rebuild from host** (Docker required â€” `feeding/regenerate` does NOT work in the Docker workflow)

```powershell
cd "c:\Return of Reckoning\Interface\WAR_API_Ref"
make build ; make generate-platform
make build-mcp
docker compose restart mcp-server
```

Done. The new knowledge is now live in `docs/war-api/meta/conventions.md`.

---

## Flow Overview

```
.feed.json files  â†’  ingest_batch (persist+auto_accept)  â†’  promote_all_accepted  â†’  make build + generate-platform  â†’  make build-mcp
```

Old flow (still valid for granular control):

```
ingest_batch (persist)  â†’  list_pending  â†’  review (per ID)  â†’  promote (per ID)  â†’  make build + generate-platform  â†’  make build-mcp
```

---

## Seed Files

| Seed file | What goes there |
|-----------|----------------|
| `docs/platform/seeds/war_api_facts.md` | Lua API functions, namespaces, runtime behaviour (e.g., `Tooltips.*`, `Cursor.Data.*`) |
| `docs/platform/seeds/xml_conventions.md` | XML element attributes, drag-drop, input routing caveats |
| `docs/platform/seeds/lua_conventions.md` | Lua tooling/encoding issues (e.g., UTF-8 BOM), scripting conventions |

All three seeds are read by `make generate-platform` and surfaced in `docs/war-api/meta/conventions.md`.

---

## Observation File Template

One `.feed.json` per fact cluster. No `.md` narrative file needed.

```json
{
  "schema_version": "feeding.observation.v1",
  "entry_id": "myaddon_short_slug",
  "status": "candidate",
  "source": {
    "addon": "MyAddon",
    "date_utc": "2026-04-14",
    "validation": "dev_observed",
    "context": "What you were doing when you found this"
  },
  "target_seeds": ["docs/platform/seeds/war_api_facts.md"],
  "confidence": { "overall": "HIGH", "rationale": "Confirmed at runtime" },
  "claims": [
    {
      "claim_id": "c1",
      "kind": "usage_pattern",
      "confidence": "HIGH",
      "statement": "Short factual statement about the engine behaviour",
      "guidance": "How addon developers should use this knowledge"
    }
  ],
  "evidence": [
    {
      "evidence_id": "e1",
      "type": "code_change",
      "summary": "Fixed the bug by doing X â€” confirmed it works in-game"
    }
  ],
  "promotion": { "notes": "Validated at runtime in AdvancedPetAssist session" }
}
```

Place the file under `docs/platform/feeding/<topic>/myslug.feed.json`.  
The batch ingest auto-discovers **all** `.feed.json` files under `docs/platform/feeding/`.

### Confidence levels

| Level | Use when |
|-------|----------|
| `HIGH` | Confirmed at runtime; fact is unambiguous |
| `MEDIUM` | Observed in one addon; reasonable but not cross-validated |
| `LOW` | Theory only; needs more testing |

### Claim kinds

| Kind | Use when |
|------|----------|
| `behavior_caveat` | Engine quirk, trap, or workaround |
| `usage_pattern` | Recommended idiom or pattern |
| `signature_hint` | Undocumented function exists and this is its shape |

---

## Writable vs Read-only Container

| Scenario | Port | Command |
|----------|------|---------|
| Normal (read-only, query only) | 8091 | `make up-mcp` |
| Feeding / promoting (writable) | 8092 | `$env:MCP_SERVER_PORT=8092; docker compose up -d --build mcp-server` |

**Only the writable container can persist observations to the queue file.** The default compose service at 8091 mounts `docs/platform/feeding` read-only.

---

## Step-by-Step (Granular Control)

### 1. Validate without persisting (dry run)

```powershell
$r = Invoke-RestMethod -Uri "http://localhost:8092/mcp" -Method Post `
  -ContentType "application/json" -Body '{"jsonrpc":"2.0","id":1,"method":"feeding/ingest_batch","params":{"dry_run":true,"continue_on_error":true}}'
$r.result | ConvertTo-Json -Depth 3
```

### 2. Persist and auto-accept

```powershell
$r = Invoke-RestMethod -Uri "http://localhost:8092/mcp" -Method Post `
  -ContentType "application/json" -Body '{"jsonrpc":"2.0","id":1,"method":"feeding/ingest_batch","params":{"persist":true,"auto_accept":true,"continue_on_error":true}}'
"Accepted: $($r.result.accepted_count) / Rejected: $($r.result.rejected_count)"
```

If you want manual review, omit `auto_accept` and use `feeding/review` per ID.

### 3. List pending (optional sanity check)

```powershell
$r = Invoke-RestMethod -Uri "http://localhost:8092/mcp" -Method Post `
  -ContentType "application/json" -Body '{"jsonrpc":"2.0","id":1,"method":"feeding/list_pending_observations","params":{"status_filter":"accepted"}}'
$r.result.observations | ForEach-Object { "$($_.observation_id): $($_.status)" }
```

### 4. Promote all accepted

```powershell
$r = Invoke-RestMethod -Uri "http://localhost:8092/mcp" -Method Post `
  -ContentType "application/json" -Body '{"jsonrpc":"2.0","id":1,"method":"feeding/promote_all_accepted","params":{}}'
"Promoted: $($r.result.promoted_count), Skipped (dup): $($r.result.skipped_count)"
```

Or promote individually: `feeding/promote` with `{ "observation_id": "my_id" }`.

### 5. Rebuild (must run from host, not MCP)

```powershell
cd "c:\Return of Reckoning\Interface\WAR_API_Ref"
make build            # rebuilds generator + MCP images
make generate-platform # runs generator container, updates docs/war-api
make build-mcp        # bakes updated docs/war-api into new MCP image
docker compose restart mcp-server
```

> **Why not `feeding/regenerate`?**  
> The MCP server runs inside Docker and does not have a `go` binary or the generator image. `feeding/regenerate` calls `go run ./tools/api_doc_gen` which will fail in this project's Docker setup. Always use `make generate-platform` from the host instead.

### 6. Verify

```powershell
Select-String -Path "docs\war-api\meta\conventions.md" -Pattern "your_keyword"
```

---

## Troubleshooting

| Problem | Cause | Fix |
|---------|-------|-----|
| `ingest_batch` returns 0 accepted | All files already in queue as promoted | Nothing to do â€” they were already promoted |
| `promote_all_accepted` returns 0 promoted | No accepted observations in queue | Run `ingest_batch` with `auto_accept` first |
| Seed content not in `conventions.md` after `generate-platform` | Seeded `lua_conventions.md` or `war_api_facts.md` but those seeds weren't being read | Fixed in generator â€” `make build` first |
| `feeding/regenerate` returns error | `go` not in MCP container | Use `make generate-platform` from host |
| `'=' expected near 'Â»'` in game | UTF-8 BOM in .lua file | Save with `[System.Text.UTF8Encoding]::new($false)` â€” see `lua_conventions.md` seed |
| Port 8092 refused | Writable container not started | `$env:MCP_SERVER_PORT=8092; docker compose up -d --build mcp-server` |

---
