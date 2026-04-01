---
applyTo: '**/*.lua'
---

# Lua Authoring Rules (WAR)

## Scope

- Keep this file to coding conventions only.
- API facts belong in:
  - `docs/platform/`
  - `docs/platform/seeds/`
  - MCP server responses

## Runtime Constraints

- Target Lua 5.0-style runtime constraints used by WAR UI.
- Use `math.mod(a, b)` rather than `%`.
- Do not use unsupported control-flow features (for example, `continue`).

## String and Type Conventions

- UI-facing strings are wstrings.
- Use `L"literal"` or `towstring(value)` for UI text.
- Keep concatenation type-safe for wstrings.

## Sparse Data Handling

- Treat inventory/buff/event payload tables as potentially sparse.
- Use `pairs()` for sparse structures.
- Do not assume dense numeric indexing unless confirmed.

## Addon Structure Conventions

- Keep one clear addon namespace table.
- Keep init/shutdown flow explicit and reversible.
- Register/unregister handlers and slash commands symmetrically.
- Prefer small wrappers around engine calls so behavior is testable/reviewable.

## Defensive Coding

- Validate window/symbol existence before use where practical.
- Use guard clauses around optional runtime state.
- Preserve confidence/uncertainty notes when behavior is inferred.

## MCP Guidance

- For API lookup, call MCP tools/resources rather than relying on memory.
- If code introduces repeated patterns or clarified behavior, propose seed updates under `docs/platform/seeds/`.
