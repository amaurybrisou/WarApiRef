o# Feeding Observation Pattern v1

## Goal

Define a normalized, machine-ingestible intake shape for platform findings that may later feed WAR_API_Ref seeds and MCP context.

## Required File Pair

For each intake document, create:

- Human narrative: `<topic>/<slug>.md`
- Machine sidecar: `<topic>/<slug>.feed.json`

Example:

- `docs/platform/feeding/xml/whohealedme_xml_input_and_layout.md`
- `docs/platform/feeding/xml/whohealedme_xml_input_and_layout.feed.json`

## Schema

- Sidecar JSON must validate against:
  - `docs/platform/feeding/model/observation.v1.schema.json`
- `schema_version` must be:
  - `feeding.observation.v1`

## Mapping Rules

- `entry_id`: stable snake_case identifier for dedupe and indexing.
- `source`: who/when/how the observation was validated.
- `target_seeds`: one or more seed docs the finding could promote into.
- `claims[]`: atomic statements, each with confidence and practical guidance.
- `evidence[]`: concrete supporting records referenced by `evidence_refs`.
- `promotion`: gate criteria and notes for seed promotion readiness.

## MCP Feed Use

The `.feed.json` files are the canonical machine feed input.

## Ingestion Interface

Use one of the following MCP interfaces:

1. Tool call: `tools/call` with `name = "ingest_observation"`
2. Direct RPC method: `feeding/ingest`
3. Batch RPC method: `feeding/ingest_batch`
4. Batch tool call: `tools/call` with `name = "ingest_observation_batch"`

Minimal payload shape:

```json
{
  "observation": {
    "schema_version": "feeding.observation.v1",
    "entry_id": "example_entry",
    "status": "candidate",
    "source": {"addon": "WhoHealedMe", "date_utc": "2026-04-01", "validation": "mixed", "context": "example"},
    "target_seeds": ["docs/platform/seeds/xml_conventions.md"],
    "confidence": {"overall": "MEDIUM", "rationale": "Implementation and user validated"},
    "claims": [{"claim_id": "c1", "kind": "usage_pattern", "confidence": "MEDIUM", "statement": "example", "guidance": "example"}],
    "evidence": [{"evidence_id": "e1", "type": "code_change", "summary": "example"}],
    "promotion": {"notes": "review before promotion"}
  },
  "dry_run": true,
  "source_path": "docs/platform/feeding/xml/example.feed.json"
}
```

Batch payload shape:

```json
{
  "dry_run": true,
  "persist": false,
  "continue_on_error": true,
  "limit": 500
}
```

Persistence behavior:

- Set `persist=true` and `dry_run=false` to append accepted records into:
  - `docs/platform/feeding/review_queue/accepted.ndjson`
- Optional `queue_file` can override the output target.

Recommended pipeline behavior:

1. Load all `docs/platform/feeding/**/*.feed.json`.
2. Validate each record against `observation.v1.schema.json`.
3. Reject invalid records with structured errors.
4. Group accepted records by `target_seeds` and confidence.
5. Emit promotion candidates into seed-review workflows.

## Confidence Policy

- Keep confidence conservative:
  - `HIGH`: strongly corroborated by canonical/multi-source evidence.
  - `MEDIUM`: user-validated or implementation-backed but not broad canonical proof.
  - `LOW`: tentative hypothesis; keep out of seed promotion by default.