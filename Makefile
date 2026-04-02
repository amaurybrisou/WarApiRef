SHELL := /bin/sh
SOURCE_ROOT := /workspace/data/raw

.PHONY: build generate-addon generate-platform generate-site preview up down build-mcp run-mcp up-mcp test-mcp

build:
	docker compose build api-doc-gen docs-site mcp-server

generate-addon:
	docker compose run --rm api-doc-gen generate addon $(SOURCE_ROOT) /workspace/docs/addon-api

generate-platform:
	docker compose run --rm api-doc-gen generate platform /workspace/docs/addon-api /workspace/docs/war-api --source-root $(SOURCE_ROOT)

generate-site:
	docker compose run --rm api-doc-gen generate site /workspace/docs/war-api /workspace/docs/site/content

preview:
	docker compose up --build docs-site

up:
	docker compose up -d --build docs-site

down:
	docker compose down

build-mcp:
	docker compose build mcp-server

run-mcp:
	docker compose run --rm mcp-server --transport stdio --docs-root /workspace/docs/war-api

up-mcp:
	docker compose up --build mcp-server

test-mcp:
	docker run --rm -v "$$(pwd):/workspace" -w /workspace golang:1.22-alpine go test ./tools/mcp_server/...