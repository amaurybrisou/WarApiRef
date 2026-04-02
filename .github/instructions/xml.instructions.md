---
applyTo: '**/*.xml'
---

# XML Authoring Rules (WAR UI)

## Scope

- API and event facts belong in:
  - `docs/war-api/`
  - `docs/platform/seeds/`
  - MCP server responses

## Layout and Naming

- Use consistent, explicit window naming and `$parent` expansion patterns.
- Keep anchor intent readable; avoid ambiguous offset chains.
- Prefer template-based repeated UI elements over duplicated blocks.

## Input and Event Discipline

- Keep `handleinput`/`skipinput` choices explicit and intentional.
- Register only required handlers; avoid accidental event fan-out.
- Do not assert event behavior without canonical documentation evidence.

## Reliability Practices

- Keep template/container responsibilities clear.
- Avoid mixing visual and interaction responsibilities in one node when possible.
- Preserve uncertainty when XML runtime behavior is inferred, not documented.

## MCP Guidance

- Resolve handler semantics and event contracts through MCP/docs sources of truth.
- If implementation confirms improved behavior knowledge, propose seed updates under `docs/platform/seeds/`.
