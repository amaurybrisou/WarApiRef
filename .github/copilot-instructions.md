# Copilot Instructions For WAR_API_Ref

## Scope

This workspace builds a multi-phase WAR addon API documentation pipeline and static docs site.

- Generator: [tools/api_doc_gen/main.go](../tools/api_doc_gen/main.go)
- MCP server: [tools/mcp_server/main.go](../tools/mcp_server/main.go)
- Generated platform docs (source of truth for MCP): [docs/war-api](../docs/war-api)
- Site content and app: [docs/site](../docs/site)

For architecture and contract context, read [README.md](../README.md) first.

## Non-Negotiable Rules

- Keep changes minimal and targeted to the request.
- Do not invent relationships, correlations, or graph edges.
- Prefer explicit evidence over heuristics.
- Preserve behavior unless the task explicitly requires behavioral change.
- Keep structural and behavioral concepts separate:
	- XML element types are not XML events.
	- XML events are not Lua functions.

If a node is dangling, classify the failure before patching:

1. no evidence exists
2. emission rule is missing
3. XML-Lua correlation is missing
4. model vocabulary is insufficient

## Phase-Based Debugging

When diagnosing parser/semantic/graph issues, trace the same symbol through phases in order:

1. XML parse
2. XML structural facts
3. Lua parse
4. XML-Lua correlation
5. semantic enrichment
6. catalog and graph emission

Do not patch graph emission if evidence is missing upstream.

## Commands

Prefer Make targets for reproducible local workflows:

- Build images: `make build`
- Generate addon docs: `make generate-addon`
- Generate platform docs: `make generate-platform`
- Generate site content: `make generate-site`
- Preview site: `make preview`
- Run MCP in stdio mode: `make run-mcp`
- Run MCP tests: `make test-mcp`

Command definitions are canonical in [Makefile](../Makefile).
Container wiring and ports are canonical in [docker-compose.yml](../docker-compose.yml).

## Validation Expectations

After meaningful changes:

1. Validate edited files compile/lint cleanly.
2. Regenerate affected outputs when touching generator semantics.
3. Compare before/after behavior for the targeted symbols.
4. Report unresolved gaps explicitly.

If full runtime validation is not possible, say so clearly.

## Docs And Knowledge Workflow

- Treat [docs/addon-api](../docs/addon-api) as intermediate generated output.
- Treat [docs/war-api](../docs/war-api) as canonical generated platform output.
- MCP tools/resources should align with generated artifacts under [docs/war-api](../docs/war-api).
- For durable new knowledge, follow the observation lifecycle documented in [README.md](../README.md).

## Change Reporting

For non-trivial tasks, summarize:

1. files changed
2. what was fixed and why
3. what was validated
4. what remains unresolved
