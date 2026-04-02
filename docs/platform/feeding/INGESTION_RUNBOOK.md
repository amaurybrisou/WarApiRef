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
