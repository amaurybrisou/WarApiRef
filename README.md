# WAR Addon API Consolidated Docs Pipeline

This repository now exposes a consolidated documentation pipeline with a single generator and a static site.

## Structure

- `tools/api_doc_gen`: Go generator (addon, platform, site modes)
- `tools/mcp_server`: MCP server over canonical `docs/war-api` artifacts
- `data/raw`: logical source root for addon scanning
- `docs/addon-api`: generated addon-layer API docs (intermediate, regenerated on demand)
- `docs/war-api`: generated WAR platform docs, semantic navigation, and graph JSON
- `docs/site`: browsable static site and generated site content
- `infra/docker`: Dockerfiles for generator and site containers

## Generator Commands

- `go run tools/api_doc_gen/main.go generate addon ./data/raw ./docs/addon-api`
- `go run tools/api_doc_gen/main.go generate platform ./docs/addon-api ./docs/war-api`
- `go run tools/api_doc_gen/main.go generate site ./docs/war-api ./docs/site/content`

## Docker-Only Workflow

- Build images:
  - `make build`
- Generate addon docs:
  - `make generate-addon`
- Generate platform docs:
  - `make generate-platform`
- Generate site content:
  - `make generate-site`
- Run preview site:
  - `make preview`

## MCP Server

The MCP server reads `docs/war-api` as its source of truth and exposes tools/resources for symbol lookup, graph navigation, confidence reasoning, and snippet scaffolding.

## MCP Configuration

The server supports two transports:

- `stdio`: for MCP clients that launch the process directly
- `http`: for network access via `/mcp`

Runtime flags:

- `--transport`: `stdio` (default) or `http`
- `--docs-root`: path to generated WAR docs root (default: `docs/war-api`)
- `--listen`: HTTP bind address when `--transport http` (default: `:8091`)

Environment variables:

- `WAR_API_ROOT`: fallback docs root used when `--docs-root` is not passed
- `MCP_SERVER_PORT`: Docker Compose host port override for HTTP mode (default: `8091`)

Examples:

- Stdio mode (recommended for local MCP client integration):
  - `make run-mcp`
- HTTP mode with Docker Compose:
  - `MCP_SERVER_PORT=8092 docker compose up --build mcp-server`
- Direct binary run over HTTP:
  - `go run ./tools/mcp_server --transport http --listen :8091 --docs-root ./docs/war-api`

Health check via MCP endpoint:

- `curl -s http://localhost:8091/mcp -H "content-type: application/json" -d '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}'`

Notes:

- The server expects generated docs to exist under the configured docs root.
- Generate/rebuild docs first when needed:
  - `make generate-addon`
  - `make generate-platform`

Contract layout:

- `tools/mcp_server/model`: shared entities, confidence, warnings, error model
- `tools/mcp_server/schema`: IDs, request/response structs, supported names, `tool_contracts.json`
- `tools/mcp_server/tools`: tool-specific handlers with stable output shapes
- `tools/mcp_server/resources`: resource resolvers with stable URI handling
- `tools/mcp_server/server`: registry, validation, dispatch, and transport adapters

- Local stdio mode (for MCP clients):
  - `make run-mcp`
- Run tests:
  - `make test-mcp`
- Build MCP container:
  - `make build-mcp`
- Run MCP server over HTTP (container):
  - `make up-mcp`

Default HTTP endpoint: `http://localhost:8091/mcp`

Supported tools:

- `lookup_symbol`
- `search_symbols`
- `get_related_symbols`
- `get_event_flow`
- `get_xml_handler_info`
- `find_usage_examples`
- `explain_confidence`
- `scaffold_addon_snippet`
- `explain_symbol_usage`

Supported resource URIs:

- `war-api://symbols/<canonical_name>`
- `war-api://events/<event_name>`
- `war-api://xml-handlers/<handler_name>`
- `war-api://patterns/<pattern_name>`
- `war-api://meta/conventions`
- `war-api://meta/inference-rules`
- `war-api://meta/coverage-report`
- `war-api://graph/<canonical_name>`
- `war-api://navigation/tree`
- `war-api://navigation/sidebar`

Machine-readable contract summary:

- `tools/mcp_server/schema/tool_contracts.json`

## Site Features

The static app in `docs/site` provides:

- full-text search from `docs/site/content/search/index.json`
- navigation tree loaded from `docs/site/content/navigation/tree.json`
- sidebar navigation paths from generated content links
- graph exploration from `docs/site/content/graph/api_graph.json`
- markdown page rendering with internal link rewrites
- MCP Live Console panel for quick `initialize` and `tools/list` checks against a running MCP endpoint

### Docs UI MCP Console

After opening the docs site, use the **MCP Live Console** panel at the top of the page:

1. Set the MCP endpoint URL (default `http://127.0.0.1:8091/mcp`)
2. Click `Initialize` to test handshake response
3. Click `Tools List` to inspect exposed tools and `inputSchema` contracts

This is a lightweight diagnostics UI for MCP health checks and client-contract verification.

## Notes

`data/raw` uses `.api_doc_gen_source_root` to point to the workspace addon corpus root. This avoids copying all addon directories while keeping the required command contract.

When running through Docker Compose, the service mounts the parent AddOns directory as `/workspace_addons` (read-only), and `data/raw/.api_doc_gen_source_root` resolves to that path.

`docs/addon-api` is intentionally treated as an intermediate output and is gitignored. Regenerate it with `make generate-addon` when needed.
