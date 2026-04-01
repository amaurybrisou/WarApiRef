# WAR API Facts Seed

## Intent

Normalized snippets of high-value platform facts for tooling and MCP context.

## Fact Style

- Keep facts atomic and verifiable.
- Prefer canonical symbol names as seen in generated docs.
- Attach confidence labels when facts are inferred.

## Seeded Baselines

- Many runtime collections are sparse in practice; iterate defensively.
- UI behavior frequently depends on XML configuration plus Lua runtime calls.
- Symbol confidence should distinguish source confidence from semantic confidence.

## Confidence Tags

- `HIGH`: directly evidenced in canonical docs and repeated usage
- `MEDIUM`: partial evidence or relation-based inference
- `LOW`: weak evidence; treat as provisional
