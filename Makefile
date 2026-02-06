# Project Variables
export PATH := $(shell go env GOPATH)/bin:$(PATH)
PROJECT_NAME := dilocash-oss
BIN_DIR := ./bin
PROTO_DIR := ./proto
GEN_DIR := ./gen
DOCS_DIR := ./docs/diagrams

# Tools
BUF := buf
SQLC := sqlc
GOVERTER := goverter
MMDC := mmdc # Mermaid CLI (requires: npm install -g @mermaid-js/mermaid-cli)

.PHONY: all help dev build generate generate-code generate-docs db-up db-down check lint test clean tidy adr jwt-decode

all: generate check

# --- Generation Suite ---

generate: generate-code generate-docs ## Run all code and documentation generation

generate-code: ## Generate Go/TS code from Proto, SQL and Mappers
	@echo "🏗️  Generating API Contracts (Buf)..."
	$(BUF) generate proto
	@echo "🗄️  Generating Database Layer (SQLC)..."
	$(SQLC) generate -f packages/database/sqlc.yaml
	@echo "🔄 Generating Model Mappers (Goverter)..."
	cd apps/api && $(GOVERTER) gen ./internal/mappers
	@echo "✅ Code generation complete."

sqlc: ## Generate Go models from SQL (sqlc)
	@echo "🗄️  Generating Database Layer (SQLC)..."
	$(SQLC) generate -f packages/database/sqlc.yaml

mappers: ## Generate type-safe model mappers (goverter)
	@echo "🔄 Generating Model Mappers (Goverter)..."
	cd apps/api && $(GOVERTER) gen ./internal/mappers

generate-docs: ## Render Mermaid diagrams (.mmd) to SVG
	@echo "🎨 Rendering Mermaid diagrams to SVG..."
	# Render the Service Architecture
	$(MMDC) -i $(DOCS_DIR)/dilocash.mmd -o $(DOCS_DIR)/dilocash.svg -t neutral -b transparent || true
	# Render the Database ERD
	$(MMDC) -i $(DOCS_DIR)/database_er.mmd -o $(DOCS_DIR)/database_er.svg -t neutral -b transparent || true
	# Render the Intent Service
	$(MMDC) -i $(DOCS_DIR)/intent_service.mmd -o $(DOCS_DIR)/intent_service.svg -t neutral -b transparent || true
	@echo "✨ Documentation rendered."

# --- Development & Build ---

dev: ## Start all applications (API, Web, Mobile) via Turborepo
	pnpm turbo run dev

build: ## Build all applications
	pnpm turbo run build

db-up: ## Start Postgres and Redis via Docker Compose
	docker-compose up -d

db-down: ## Stop local infrastructure
	docker-compose down

# --- Quality Control ---

check: lint test license-check ## Run all quality checks

license-apply: ## Apply license headers to all source files
	@echo "⚖️  Applying license headers..."
	@~/go/bin/addlicense -f .license_header -v .

license-check: ## Check if source files are missing license headers
	@echo "🔍 Checking license headers..."
	@~/go/bin/addlicense -f .license_header -check .

lint: ## Run linters for Go and Protobuf
	@echo "Checking project health for $(PROJECT_NAME)..."
	go vet ./...
	$(BUF) lint

test: ## Run Go tests
	go test -v -race ./...

# --- Database Migrations ---

migrate-apply: ## Apply pending migrations to local DB
	@echo "🚀 Applying migrations..."
	cd apps/api && set -a && . ./.env && set +a && atlas migrate apply --env local

migrate-new: ## Generate a new migration file (usage: make migrate-new name=add_users)
	@echo "📝 Generating new migration: $(name)..."
	cd apps/api && set -a && . ./.env && set +a && atlas migrate diff $(name) --env local

# --- Integrations & Debugging ---

ngrok: ## Expose local API for WhatsApp/Telegram webhooks
	ngrok http 8080

bot-test: ## Send a mock voice payload to the intent engine
	curl -X POST http://localhost:8080/v1/adapters/test -F "audio=@test.m4a"

jwt-decode: ## Decode a JWT token for debugging (requires jq)
	@echo $(token) | cut -d. -f2 | base64 --decode | jq

# --- Documentation ---

adr: ## Scaffold a new ADR (usage: make adr n=0005 t="use_redis_cache")
	@NUMBER=$(n); \
	TITLE=$(t); \
	cp docs/adr/0000-template.md docs/adr/$$NUMBER-$$TITLE.md; \
	echo "Created ADR: docs/adr/$$NUMBER-$$TITLE.md"

# --- Cleanup ---

clean: ## Remove generated binaries and code
	rm -rf $(BIN_DIR)
	rm -rf $(GEN_DIR)/go/*
	rm -rf $(GEN_DIR)/ts/*
	@echo "🧹 Cleaned all generated assets."

tidy: ## Tidy Go modules
	cd apps/api && go mod tidy
	cd infra && go mod tidy

# --- Help ---

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help