# Project Variables
export PATH := $(shell go env GOPATH)/bin:$(PATH)
PROJECT_NAME := dilocash-oss
BIN_DIR := ./bin
PROTO_DIR := ./packages/proto
GEN_DIR := ./gen
DOCS_DIR := ./docs/diagrams
LICENSE_IGNORE := -ignore "apps/api/migrations/**" -ignore "node_modules/**" -ignore "apps/web/node_modules/**" -ignore "apps/web/.next/**" -ignore "apps/web/next-env.d.ts" -ignore "pnpm-lock.yaml"

# Tools
BUF := buf
SQLC := sqlc
GOVERTER := goverter
# MMDC := mmdc # Mermaid CLI (requires: npm install -g @mermaid-js/mermaid-cli)

.PHONY: all help dev build generate generate-code generate-docs db-up db-down check lint test clean tidy adr jwt-decode

all: generate check

# --- Generation Suite ---

generate: generate-code generate-docs ## Run all code and documentation generation

generate-code: ## Generate Go/TS code from Proto, SQL and Mappers
	@echo "🏗️  Generating API Contracts (Buf)..."
	@cd $(PROTO_DIR) && $(BUF) generate
	@echo "🗄️  Generating Database Layer (SQLC)..."
	@$(SQLC) generate -f packages/database/sqlc.yaml
	@echo "🔄 Generating Model Mappers (Goverter)..."
	@cd apps/api && $(GOVERTER) gen ./internal/mappers
	@mkdir -p apps/api/internal/generated/mappers
	@mv apps/api/internal/mappers/generated/generated.go apps/api/internal/generated/mappers/
	@rm -rf apps/api/internal/mappers/generated
	@echo "✅ Code generation complete."

sqlc: ## Generate Go models from SQL (sqlc)
	@echo "🗄️  Generating Database Layer (SQLC)..."
	$(SQLC) generate -f packages/database/sqlc.yaml

generate-docs: ## Render Mermaid diagrams (.mmd) to SVG
# 	@echo "🎨 Rendering Mermaid diagrams to SVG..."
# 	# Render the Service Architecture
# 	$(MMDC) -i $(DOCS_DIR)/dilocash.mmd -o $(DOCS_DIR)/dilocash.svg -t neutral -b transparent || true
# 	# Render the Database ERD
# 	$(MMDC) -i $(DOCS_DIR)/database_er.mmd -o $(DOCS_DIR)/database_er.svg -t neutral -b transparent || true
# 	# Render the Intent Service
# 	$(MMDC) -i $(DOCS_DIR)/intent_service.mmd -o $(DOCS_DIR)/intent_service.svg -t neutral -b transparent || true
# 	@echo "✨ Documentation rendered."

# --- Development & Build ---

install: ## Install all dependencies at the root
	pnpm install

dev: install ## Start all applications (API, Web, Mobile) via Turborepo
	pnpm dev

dev-mobile: install ## Start Mobile Android app
	cd apps/mobile && pnpm run android

dev-web: install ## Start Web app via Turborepo
	pnpm dev --filter @dilocash/web

dev-web-api: install ## Start Web app and API via Turborepo
	pnpm dev --filter @dilocash/web --filter @dilocash/api

dev-api: install ## Start API via Turborepo
	pnpm dev --filter @dilocash/api

supabase-up:
	@echo "🔐 Starting Supabase..."
	cd supabase && npx supabase start

supabase-down:
	@echo "🔐 Stopping Supabase..."
	cd supabase && npx supabase stop

build: ## Build all applications via Turborepo
	@echo "🏗️  Building all applications..."
	pnpm build

build-api: ## Build API binary directly (faster than turbo build)
	@echo "🏗️  Building API..."
	cd apps/api && go build -o bin/api cmd/main.go
	@echo "✅ API built to apps/api/bin/api"

db-up: ## Start Postgres and Redis via Docker Compose
	docker-compose up -d

db-down: ## Stop local infrastructure
	docker-compose down

# --- Quality Control ---

check: lint test license-check ## Run all quality checks

license-apply: ## Apply license headers to all source files
	@echo "⚖️  Applying license headers..."
	@~/go/bin/addlicense -f .license_header -v $(LICENSE_IGNORE) .

license-check: ## Check if source files are missing license headers
	@echo "🔍 Checking license headers..."
	@~/go/bin/addlicense -f .license_header -check $(LICENSE_IGNORE) .

lint: ## Run linters for Go and Protobuf
	@echo "🏥 Checking project health for $(PROJECT_NAME)..."
	@echo "🔍 Linting API modules..."
	cd apps/api && go vet ./...
	@echo "🔍 Linting Infrastructure modules..."
	cd infra && go vet ./...
	@echo "🧹 Linting Protobuf definitions..."
	cd $(PROTO_DIR) && $(BUF) lint
	@echo "✨ All checks passed!"

test: ## Run Go tests
	@echo "🚀 Running tests..."
	cd apps/api && go test -v -race ./...

# --- Database Migrations ---

migrate-apply: ## Apply pending migrations to local DB
	@echo "🚀 Applying migrations..."
	cd apps/api && set -a && . ./.env && set +a && atlas migrate apply --env local --allow-dirty

migrate-hash: ## Generate a new migration file (usage: make migrate-new name=add_users)
	@echo "📝 Generating migration hash..."
	cd apps/api && set -a && . ./.env && set +a && atlas migrate hash --env local

migrate-new: ## Generate a new migration file (usage: make migrate-new name=add_users)
	@echo "📝 Generating new migration: $(name)..."
	cd apps/api && set -a && . ./.env && set +a && atlas migrate diff $(name) --env local

# --- Integrations & Debugging ---

bot-test: ## Send a mock voice payload to the intent engine
	curl -X POST http://localhost:8080/v1/adapters/test -F "audio=@test.m4a"

jwt-decode: ## Decode a JWT token for debugging (usage: make jwt-decode token=<JWT>)
	@if [ -z "$(token)" ]; then \
		echo "❌ Error: No token provided"; \
		echo "Usage: make jwt-decode token=<YOUR_JWT_TOKEN>"; \
		exit 1; \
	fi
	@echo $(token) | cut -d. -f2 | base64 --decode | jq

# --- Documentation ---

adr: ## Scaffold a new ADR (usage: make adr n=0005 t="use_redis_cache")
	@NUMBER=$(n); \
	TITLE=$(t); \
	cp docs/adr/0000-template.md docs/adr/$$NUMBER-$$TITLE.md; \
	echo "Created ADR: docs/adr/$$NUMBER-$$TITLE.md"

# --- Cleanup ---

clean: clean-ui ## Remove generated binaries and code
	rm -rf $(BIN_DIR)
	rm -rf node_modules
	rm -rf apps/api/bin
	rm -rf .turbo
	@echo "🧹 Cleaned all generated assets."

clean-full: clean ## Remove generated binaries and code
	rm -rf pnpm-lock.yaml
	@echo "🧹 Cleaned all generated assets."

clean-ui: ## Remove generated ui code
	rm -rf apps/web/node_modules apps/web/.next apps/web/public/sw*
	rm -rf apps/mobile/node_modules apps/mobile/.expo apps/mobile/android apps/mobile/ios
	rm -rf packages/ui/node_modules
	@echo "🧹 Cleaned all generated ui assets."

clean-expo: ## clears expo
	cd apps/mobile && pnpm run clear

tidy: ## Tidy Go modules
	cd apps/api && go mod tidy
	cd infra && go mod tidy

# --- Help ---

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help