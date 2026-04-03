# Feeding Runbook

This runbook explains how to feed observations into the WAR_API_Ref model ingestion flow.

## What "feeding the model" means here

- Input records live as `*.feed.json` files under `docs/platform/feeding/**`.
- The ingestion endpoint validates records against the normalized observation pattern.
- Accepted records can be persisted into the review queue (`accepted.ndjson`) for later seed-promotion workflows.

## Required inputs

1. A human narrative file:
   - `docs/platform/feeding/<topic>/<slug>.md`
2. A machine sidecar file:
   - `docs/platform/feeding/<topic>/<slug>.feed.json`
3. Sidecar schema version:
   - `schema_version: "feeding.observation.v1"`

Reference pattern:
- `docs/platform/feeding/model/observation_pattern.v1.md`
- `docs/platform/feeding/model/observation.v1.schema.json`

## Safety first: backup before ingest

PowerShell example from repository root:

```powershell
$ts = Get-Date -Format "yyyyMMdd_HHmmss"
$backupRoot = Join-Path "docs\platform\feeding\backup" $ts
New-Item -ItemType Directory -Path $backupRoot -Force | Out-Null
Copy-Item "docs\platform\feeding\xml\whohealedme_xml_input_and_layout.feed.json" (Join-Path $backupRoot "whohealedme_xml_input_and_layout.feed.json.bak") -Force
if (Test-Path "docs\platform\feeding\review_queue") {
  Copy-Item "docs\platform\feeding\review_queue" (Join-Path $backupRoot "review_queue") -Recurse -Force
}
```

## Start MCP ingestion endpoint

### Option A: local binary

Use this only if Go is installed.

```powershell
go run .\tools\mcp_server\main.go --transport http --listen :8091 --docs-root docs\war-api --feeding-root docs\platform\feeding
```

### Option B: Docker compose

```powershell
docker compose up --build -d mcp-server
```

Note: the compose service currently mounts `/workspace` read-only, so persistence writes can fail in that mode.

## Dry-run ingest (recommended first)

Dry-run validates without writing queue output.

```powershell
$payload = @{
  jsonrpc = "2.0"
  id = 1
  method = "feeding/ingest_batch"
  params = @{ dry_run = $true; continue_on_error = $true }
} | ConvertTo-Json -Depth 8

Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:8091/mcp" -ContentType "application/json" -Body $payload
```

Expected result includes:
- `processed_files`
- `accepted_count`
- `rejected_count`
- per-file `results[]`

## Persist accepted records

Set `dry_run=false` and `persist=true`.

```powershell
$payload = @{
  jsonrpc = "2.0"
  id = 2
  method = "feeding/ingest_batch"
  params = @{
    dry_run = $false
    persist = $true
    continue_on_error = $true
    queue_file = "/workspace/docs/platform/feeding/review_queue/accepted.ndjson"
  }
} | ConvertTo-Json -Depth 8

Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:8091/mcp" -ContentType "application/json" -Body $payload
```

If you see a read-only filesystem error, run a writable container (example):

```powershell
docker run -d --name war-api-mcp-server-rw -p 8092:8091 -v "${PWD}:/workspace" ror-mcp-server:local --transport http --listen :8091 --docs-root /workspace/docs/war-api --feeding-root /workspace/docs/platform/feeding
```

Then call `http://127.0.0.1:8092/mcp` for the persist request.

## Verify ingestion happened

1. Queue file exists:
   - `docs/platform/feeding/review_queue/accepted.ndjson`
2. Queue contains an object for your `entry_id`.
3. `source_path` points to the ingested `*.feed.json` file.

## Single-record ingest

You can ingest one record directly with method `feeding/ingest` or tool `ingest_observation`.

Minimal method payload shape:

```json
{
  "jsonrpc": "2.0",
  "id": 3,
  "method": "feeding/ingest",
  "params": {
    "observation": { "schema_version": "feeding.observation.v1", "entry_id": "example", "status": "candidate", "source": { "addon": "X", "date_utc": "2026-04-01", "validation": "mixed", "context": "example" }, "target_seeds": ["docs/platform/seeds/xml_conventions.md"], "confidence": { "overall": "MEDIUM", "rationale": "example" }, "claims": [{ "claim_id": "c1", "kind": "usage_pattern", "confidence": "MEDIUM", "statement": "example", "guidance": "example" }], "evidence": [{ "evidence_id": "e1", "type": "code_change", "summary": "example" }], "promotion": { "notes": "review" } },
    "dry_run": true,
    "persist": false,
    "source_path": "docs/platform/feeding/xml/example.feed.json"
  }
}
```

## Operational checklist

1. Create/update `.feed.json` input.
2. Backup feeding input + queue.
3. Run dry-run batch ingest.
4. Fix any rejected records.
5. Run persist batch ingest.
6. Confirm queue entry in `accepted.ndjson`.
7. Stop temporary containers/services.

---

## Review, Promote, and Regenerate Workflow

After observations are ingested into the queue, they proceed through a governed review→promotion→regeneration cycle.

### 1. Review Stage: Accept or Reject

Use `feeding/review_observation` to apply a verdict and optional reviewer commentary.

```powershell
$payload = @{
  jsonrpc = "2.0"
  id = 10
  method = "feeding/review_observation"
  params = @{
    observation_id = "whohealedme_xml_input_and_layout"
    verdict = "accept"  # or "reject"
    reviewer = "addon-maintainer@example.com"
    notes = "validated against current addon implementation; all claims corroborated"
  }
} | ConvertTo-Json -Depth 8

Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:8091/mcp" -ContentType "application/json" -Body $payload
```

**Outcome:**
- `verdict = "accept"` → observation status changes to `reviewed` (eligible for promotion)
- `verdict = "reject"` → observation status becomes `rejected` and is appended to `rejected.ndjson` durable store

All review verdicts (accept/reject) include:
- Reviewer identity
- Decision timestamp (RFC3339)
- Optional notes for decision rationale

Rejected observations are:
- Immutable (cannot be re-reviewed or promoted)
- Archived in `docs/platform/feeding/review_queue/rejected.ndjson` for audit trail
- Queryable via `feeding/list_rejected_observations`

### 2. Promotion Stage: Move to Seed Files

Once an observation reaches `reviewed` status, promote it to target seed files using `feeding/promote_observation`.

```powershell
$payload = @{
  jsonrpc = "2.0"
  id = 11
  method = "feeding/promote_observation"
  params = @{
    observation_id = "whohealedme_xml_input_and_layout"
    dry_run = $true  # Preview first
  }
} | ConvertTo-Json -Depth 8

Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:8091/mcp" -ContentType "application/json" -Body $payload
```

**Promotion Steps:**

1. **Symbol Validation** (non-blocking, advisory):
   - For each symbol_id in claims[].symbols[], attempts to resolve against generated docs
   - Issues warnings if symbols cannot be found
   - Does NOT prevent promotion (symbols may be forward-references or resolve after next regen)
   - Rationale: Observations contain informational symbol references, not canonical symbol definitions

2. **Seed Path Resolution:**
   - Reads observation's `target_seeds[]` field
   - OR uses `seed_path_override` parameter if provided
   - Validates paths don't escape workspace root (path traversal prevention)

3. **Duplicate Prevention:**
   - Checks each seed file for existing promotion marker (HTML comment)
   - Entry_id must be unique per seed file
   - Prevents accidental re-promotion of same observation

4. **Markdown Fragment Creation:**
   - Builds HTML metadata comment: `<!-- OBSERVATION:entry_id (promoted:timestamp) -->`
   - Includes: source addon, confidence level, promotion date
   - Lists all claims with confidence labels and guidance
   - Appends to target seed file(s)

5. **Status Update:**
   - Observation marked as `promoted` with timestamp
   - Queue record becomes immutable (no further review/rejection allowed)

**Example Response:**
```json
{
  "observation_id": "whohealedme_xml_input_and_layout",
  "target_seeds": ["docs/platform/seeds/xml_conventions.md"],
  "seed_updates": [
    {
      "seed_path": "docs/platform/seeds/xml_conventions.md",
      "appended": true,
      "duplicate": false
    }
  ],
  "promoted": true,
  "warnings": [
    {
      "code": "unresolved_symbol",
      "message": "symbol 'hypothetical_future_symbol' could not be resolved..."
    }
  ]
}
```

### 3. Regeneration Policy

After promotion completes, the generated docs must be rebuilt to reflect seed file updates.

**Regeneration Timing:**
- **Immediate (recommended):** Call `feeding/regenerate_from_promoted_knowledge` immediately after successful promotion
- **Batch (alternative):** Schedule nightly regeneration of all platform docs
- **Manual:** Run `make generate-platform` and `make generate-site` in repository root

**Regeneration Scope:**
- `scope = "platform"`: Rebuild `docs/war-api/` from `docs/addon-api/` + seed docs
- `scope = "site"`: Rebuild `docs/site/content/` from `docs/war-api/`
- `scope = "full"` (default): Run both platform and site regeneration

**Workflow Example (Promote + Immediate Regen):**

```powershell
# Step 1: Promote observation (dry-run first)
$promotePayload = @{
  jsonrpc = "2.0"
  id = 11
  method = "feeding/promote_observation"
  params = @{ observation_id = "whohealedme_xml_input_and_layout"; dry_run = $false }
} | ConvertTo-Json -Depth 8

$promoteResult = Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:8091/mcp" -ContentType "application/json" -Body $promotePayload
if ($promoteResult.result.promoted -eq $true) {
  Write-Host "Promotion succeeded; regenerating docs..."

  # Step 2: Regenerate (dry-run first to verify)
  $regenPayload = @{
    jsonrpc = "2.0"
    id = 12
    method = "feeding/regenerate_from_promoted_knowledge"
    params = @{ scope = "full"; dry_run = $false }
  } | ConvertTo-Json -Depth 8

  $regenResult = Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:8091/mcp" -ContentType "application/json" -Body $regenPayload
  $regenResult.result.steps | Format-Table -AutoSize
}
```

**Regeneration Output:**
```json
{
  "scope": "full",
  "success": true,
  "steps": [
    {
      "label": "generate platform docs",
      "command": "go run ./tools/api_doc_gen generate platform ./docs/addon-api ./docs/war-api",
      "output": "...",
      "success": true
    },
    {
      "label": "generate site content",
      "command": "go run ./tools/api_doc_gen generate site ./docs/war-api ./docs/site/content",
      "output": "...",
      "success": true
    }
  ]
}
```

---

## Seed File Traceability

Promoted observations include HTML metadata comments in seed files for audit trail and versioning.

**Comment Format:**
```markdown
<!-- OBSERVATION:entry_id_slug (promoted:2026-04-02T14:30:00Z) -->
> Source: AddonName | Confidence: MEDIUM | Promoted: 2026-04-02

- `MEDIUM`: Claim statement here
  - Guidance: How to apply this claim
```

**Traceability Benefits:**
- Identify which observations contributed to each seed fact
- Match seed content back to original observation and evidence
- Detect stale/superseded observations during seed maintenance
- Audit feed-back loop impact on canonical docs

**Inspecting Seed Content:**
```bash
# View all promotions in a seed file
grep "<!-- OBSERVATION:" docs/platform/seeds/xml_conventions.md

# Count observations promoted to a seed
grep -c "<!-- OBSERVATION:" docs/platform/seeds/xml_conventions.md
```

---

## Semantic Contracts: Observations Are Not Sources of Truth

### Key Principle

Feeding observations are **reviewed knowledge input only**, never the semantic source of truth. Semantic contracts remain:
- `xml-tree`: XML structural analysis (parser-driven)
- `lua-analysis`: Lua semantic parsing and scoping (parser-driven)
- `xml-lua-links`: Cross-validation between XML and Lua (generator-driven)

Observations inform seeds, seeds inform docs, docs are generated sources of truth.

**Flow:**
```
Addon Code
    ↓
Observation (evidence-backed claim: "ListData binds ListBox to Lua table")
    ↓
Review & Approval
    ↓
Seed File (formalized as reusable guidance)
    ↓
Doc Generation (integrates seed content into canonical docs)
    ↓
Generated Docs (source of truth for MCP queries)
```

### Candidate Observations Do NOT Affect Semantics Automatically

- Observations in `candidate` status are NOT visible to MCP queries
- They do NOT contribute to docs until `reviewed` → `promoted`
- They do NOT trigger regeneration on ingestion
- Manual review step enforces governance

This prevents untrusted or provisional findings from polluting semantic contracts.

---

## Troubleshooting

### Symbol Validation Warnings

**Scenario:** `unresolved_symbol` warning during promotion

**Cause:** Referenced symbol not found in generated docs

**Action:**
- Verify symbol spelling matches canonical doc names
- Check if symbol should exist: search `docs/war-api/` for similar names
- If symbol is new/missing, consider whether canonical docs need update first
- Promotion proceeds anyway; warnings are advisory

### Regeneration Failures

**Scenario:** `feeding/regenerate_from_promoted_knowledge` returns errors

**Action:**
1. Verify workspace root validation: check Makefile, docker-compose.yml, tools/api_doc_gen exist
2. Run regeneration manually: `make generate-platform` or `make generate-site` from repo root
3. Check tool output for specific errors (schema issues, missing dependencies, etc.)
4. Review recent seed file edits for syntax errors

### Queue Conflicts

**Scenario:** Observation appears in both `accepted.ndjson` and `rejected.ndjson`

**Action:**
- This should not happen (rejection excludes from accepted queue)
- If found, manually inspect files and remove duplicate entry
- Contact workflow owner to investigate root cause

