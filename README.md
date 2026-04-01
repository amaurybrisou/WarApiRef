# WAR Addon API Consolidated Docs Pipeline

This repository now exposes a consolidated documentation pipeline with a single generator and a static site.

## Structure

- `tools/api_doc_gen`: Go generator (addon, platform, site modes)
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

## Site Features

The static app in `docs/site` provides:

- full-text search from `docs/site/content/search/index.json`
- navigation tree loaded from `docs/site/content/navigation/tree.json`
- sidebar navigation paths from generated content links
- graph exploration from `docs/site/content/graph/api_graph.json`
- markdown page rendering with internal link rewrites

## Notes

`data/raw` uses `.api_doc_gen_source_root` to point to the workspace addon corpus root. This avoids copying all addon directories while keeping the required command contract.

When running through Docker Compose, the service mounts the parent AddOns directory as `/workspace_addons` (read-only), and `data/raw/.api_doc_gen_source_root` resolves to that path.

`docs/addon-api` is intentionally treated as an intermediate output and is gitignored. Regenerate it with `make generate-addon` when needed.
