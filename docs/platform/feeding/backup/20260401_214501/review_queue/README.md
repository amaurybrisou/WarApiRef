# Feeding Review Queue

This folder stores accepted ingestion records before promotion into seed docs.

## Default Queue File

- `accepted.ndjson`

Each line is one JSON object with:

- `ingested_at_utc`
- `source_path`
- `observation`

## Notes

- This queue is append-only.
- Promotion workflows should consume from this queue and create reviewed seed candidates.
- Use `dry_run=true` during validation-only runs to avoid queue writes.