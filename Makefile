SHELL := /bin/sh

.PHONY: build generate-addon generate-platform generate-site preview up down

build:
	docker compose build api-doc-gen docs-site

generate-addon:
	docker compose run --rm api-doc-gen generate addon /workspace/data/raw /workspace/docs/addon-api

generate-platform:
	docker compose run --rm api-doc-gen generate platform /workspace/docs/addon-api /workspace/docs/war-api

generate-site:
	docker compose run --rm api-doc-gen generate site /workspace/docs/war-api /workspace/docs/site/content

preview:
	docker compose up --build docs-site

up:
	docker compose up -d --build docs-site

down:
	docker compose down
