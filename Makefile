# TODO clean up Makefile

# Project Variables
PROJECT_NAME := dilocash-oss
BIN_DIR := ./bin
PROTO_DIR := ./proto
GEN_DIR := ./gen
DOCS_DIR := ./docs/diagrams

# Tools
BUF := buf
SQLC := sqlc
MMDC := mmdc # Mermaid CLI (requires: npm install -g @mermaid-js/mermaid-cli)

.PHONY: all help generate generate-code generate-docs db-up db-down check test clean

all: generate check

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'




# --- Generation Suite ---

generate: generate-code generate-docs ## Run all code and documentation generation

generate-code: ## Generate Go/TS code from Proto and SQL
	@echo "🏗️  Generating API Contracts (Buf)..."
	$(BUF) generate
	@echo "🗄️  Generating Database Layer (SQLC)..."
	$(SQLC) generate
	@echo "✅ Code generation complete."



generate-docs: ## Regenerate Mermaid diagrams and render to SVG
	@echo "🎨 Rendering Mermaid diagrams to SVG..."
	# Render the Service Architecture
	$(MMDC) -i $(DOCS_DIR)/dilocash.mmd -o $(DOCS_DIR)/dilocash.svg -t neutral -b transparent
	# Render the Database ERD
	# Note: This assumes you have the .mmd source in the docs folder
	$(MMDC) -i $(DOCS_DIR)/database_er.mmd -o $(DOCS_DIR)/database_er.svg -t neutral -b transparent
	@echo "✨ Documentation rendered."

# --- Development Infrastructure ---

db-up: ## Start Postgres and Redis via Docker Compose
	docker-compose up -d

db-down: ## Stop local infrastructure
	docker-compose down

# --- Quality Control ---

check: lint test ## Run all quality checks

lint: ## Run linters for Go and Protobuf
	go vet ./...
	$(BUF) lint

test: ## Run Go tests
	go test -v -race ./...

# --- AI Context Helper ---

context: ## Export project context for AI assistants (Antigravity)
	@chmod +x ./scripts/ai-context.sh
	@./scripts/ai-context.sh

# --- Cleanup ---

clean: ## Remove generated binaries and code
	rm -rf $(BIN_DIR)
	rm -rf $(GEN_DIR)/go/*
	rm -rf $(GEN_DIR)/ts/*
	@echo "🧹 Cleaned all generated assets."




# --- Repository Management ---

REPO_NAME := dilocash-oss

check: lint test ## Run all checks (Lint, Tests, Buf)
	@echo "Checking project health for $(REPO_NAME)..."
	go vet ./...
	staticcheck ./...
	buf lint

clean-gen: ## Remove generated code to force a fresh sync
	rm -rf gen/go/*
	rm -rf gen/ts/*

# Dilocash Monorepo Management
.PHONY: dev build db-up db-down migrate-new migrate-apply adr help

# --- Development ---
dev: ## Start all applications (API, Web, Mobile) via Turborepo
	pnpm turbo run dev

build: ## Build all applications
	pnpm turbo run build

# --- Infrastructure ---
db-up: ## Spin up PostgreSQL and pgAdmin
	docker-compose up -d

db-down: ## Shut down infrastructure
	docker-compose down

# --- Database (Atlas + SQLC) ---
migrate-apply: ## Apply pending migrations to local DB
	cd apps/api && atlas migrate apply --env local

migrate-new: ## Generate a new migration file (usage: make migrate-new name=add_users)
	cd apps/api && atlas migrate diff $(name) --env local

sqlc: ## Generate type-safe Go code from SQL queries
	cd apps/api && sqlc generate

# --- Integrations ---
ngrok: ## Expose local API for WhatsApp/Telegram webhooks
	ngrok http 8080

bot-test: ## Send a mock voice payload to the intent engine
	curl -X POST http://localhost:8080/v1/adapters/test -F "audio=@test.m4a"

# --- Documentation ---
adr: ## Scaffold a new ADR (usage: make adr n=0005 t="use_redis_cache")
	@NUMBER=$(n); \
	TITLE=$(t); \
	cp docs/adr/0000-template.md docs/adr/$$NUMBER-$$TITLE.md; \
	echo "Created ADR: docs/adr/$$NUMBER-$$TITLE.md"

jwt-decode: ## Decode a JWT token for debugging (requires jq)
	@echo $(token) | cut -d. -f2 | base64 --decode | jq

# --- Code Generation Pipeline ---
generate: ## Generate everything from Design/Contracts
	make sqlc          # From DB Schema (Mermaid -> SQL)
	buf generate       # From API Contracts (Mermaid -> Proto)
	make generate-docs # Update Diagrams from code (Optional)

# Create a new service (usage: make service name=reports)
service: 
	@echo "Defining new service: $(name)"
	# 1. Update Mermaid diagram in docs/
	# 2. Add service to proto/dilocash/v1/$(name).proto
	# 3. Run make generate

# --- Design to Proto ---
mmd-to-proto: ## Convert Mermaid diagrams to Protobuf contracts
	@echo "Parsing Mermaid diagrams..."
	python3 scripts/mermaid_to_proto.py ./docs/service.mmd ./proto/v1/api.proto
	buf generate

# --- Contract-First Sync ---
sync-design: ## Update Go code and Mermaid diagrams from Proto files
	buf generate
	# Optional: Script to inject generated mermaid into README.md
	./scripts/inject-mermaid-docs.sh 
	@echo "✅ Interfaces and Diagrams are now in sync."

# --- Documentation Sync ---
generate-docs: ## Generate .mmd files from the Source of Truth
	# Generate ER diagram from SQL
	sql-to-mermaid -i apps/api/db/schema.sql -o docs/diagrams/database_er.mmd
	# Generate Service diagrams from Proto
	protoc --mermaid_out=docs/diagrams/ proto/v1/*.proto
	@echo "✨ Documentation files updated in docs/diagrams/"

# --- Documentation Pipeline ---
gen-docs: ## Regenerate all visual assets as SVGs
	# Generate ER diagram SVG
	mmdc -i docs/diagrams/database_er.mmd -o docs/diagrams/database_er.svg -t neutral
	
	# Generate Intent Service SVG
	mmdc -i docs/diagrams/intent_service.mmd -o docs/diagrams/intent_service.svg -t neutral
	
	@echo "🎨 SVG Diagrams updated."


# --- Tidy ---
tidy:
	go mod tidy
	go mod vendor

# --- Help ---
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help