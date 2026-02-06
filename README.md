# 💸 Dilocash (Open Core)

### _Dilo y regístralo. Say it and track it._

**Dilocash** is a senior-engineered financial engine that bridges the gap between daily life and financial awareness. It is a universal app (Web + Android) that allows you to manage your economy through standard menus or by simply "talking" to your ledger.

Dilocash-oss is the community-driven engine of the Dilocash ecosystem. It provides the NLU pipeline, OIDC authentication, and the core Go API for recording transactions via voice and text.

## 🌟 Key Features

- **🎙️ Voice-to-Ledger (NLU):** Don't type your expenses. Just say: _"Dilocash, gasté 20 dólares en café con la tarjeta"_ and the app will categorize and log it instantly.
- **📊 Unified Finance Dashboard:** High-performance charts (Recharts) that work on both Next.js and React Native.
- **🏦 Multi-Currency Engine:** Native support for USD and LatAm currencies (MXN, COP, ARS) with real-time conversion.
- **🛡️ Senior-Grade Security:** WebAuthn for browser login and Biometrics (Fingerprint/FaceID) for Android.

---

## 🏗️ System Architecture

Dilocash is built with a **Staff-level** focus on precision, security, and scalability.

- **Backend:** Golang (1.22+) following **Clean Architecture**.
- **API:** gRPC + Protobuf
- **Database:** PostgreSQL with `NUMERIC` types for 100% monetary precision.
- **Cache:** Redis
- **Type Safety**: `sqlc` for Go code generation, **Goverter** for type-safe model mapping, and `Zod` for frontend validation.
- **Migrations:** **Atlas** (Declarative migration management).
- **Frontend:** **Next.js 15** (Web) and **Expo** (Mobile) via a **Turborepo** monorepo.
- **Observability:** Structured JSON logging via `slog` with request tracing.
- **Shared:** **Solito** for universal routing and **Tamagui** for shared UI components.
- **AI Engine:** OpenAI **Whisper** (Transcription) + **GPT-4o-mini** (Intent Extraction).
- **Infrastructure:** Docker Compose (Local) / Infrastructure as Code (IaC) via **Pulumi** (Go) (AWS Cloud)

## 🏗️ System Architecture

Dilocash uses a **Design-First** approach where documentation is derived from source code.

### Entity Relationship (Database)

![Database ER Diagram](./docs/diagrams/database_er.svg)

### API Service Contract

![Intent Service Diagram](./docs/diagrams/intent_service.svg)

> 💡 _To update these, modify the `.proto` or `.sql` files; GitHub Actions will handle the rest._

---

## 🚀 Local Development Setup

Follow these steps to get the Dilocash ecosystem running on your machine.

### 1. Prerequisites

Ensure you have the following installed:

- **Go** (1.22+)
- **Node.js** (20+) & **pnpm** (`sudo npm install -g pnpm`)
- **Docker** & **Docker Compose**
- **Atlas CLI** (`curl -sSf https://atlasgo.sh | sh`)
- **sqlc** (`go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`)
- **goverter** (`go install github.com/jmattheis/goverter/cmd/goverter@latest`)
- **Buf CLI** (`brew install bufbuild/buf/buf` or see [buf.build](https://buf.build/docs/installation))
- **protoc-gen-go** (`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`)
- **protoc-gen-go-grpc** (`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`)
- **mmdc** (`npm install -g @mermaid-js/mermaid-cli`)
- **Pulumi** CLI (for cloud deployment)

### 2. Clone and Install

```bash
git clone https://github.com/dilocash/dilocash-oss.git
cd dilocash-oss
```

### 3. Infrastructure & Database

Spin up the database and apply the initial schema migrations.

```bash
# Start Postgres and pgAdmin
make db-up

# Apply migrations to the local database
make migrate-apply

# Generate Go models, API contracts, and model mappers
make generate-code

# Refresh SVG diagrams (requires mmdc)
make generate-docs
```

### 4. Environment Variables

Copy the environment file: `cp .env.example .env` file in `apps/api/` and `apps/web/` and edit its properties according to your environment.

### 5. Running the Apps

Use the **Makefile** to launch the entire monorepo (API, Web, and Mobile) simultaneously.

```bash
make dev

```

- **Web Dashboard:** `http://localhost:3000`
- **Go API:** `http://localhost:8080`
- **pgAdmin:** `http://localhost:5050` (Login: `admin@dilocash.com` / `admin`)

---

## 🎙️ The Voice-to-Ledger Flow

To test the voice feature locally:

1. Open the mobile app (via Expo Go).
2. Press and hold the **"Dilocash"** mic button.
3. Speak: _"Gasté cincuenta pesos en tacos"_ (I spent 50 pesos on tacos).
4. The Go backend will transcribe and parse the intent, returning a confirmation card to the UI.

## 🔗 Universal Entry Points

Dilocash isn't just an app; it's a financial layer that lives where you do.

- **📱 WhatsApp/Telegram:** Send a voice note or text to your private Dilocash Bot. It uses the same NLU engine to log your data.
- **⌚ WearOS (Android Watch):** A "Quick-Log" interface for the wrist. Tap, speak, and confirm.

### High-Level Flow

All entry points funnel into the `IntentService`:
`Adapter (Telegram) -> UseCase (ParseIntent) -> NLU (GPT-4o) -> Repo (Postgres)`

---

## 📂 Project Structure

```text
.
├── .github/
│   └── workflows/
│       └── docs-sync.yml        # ADR-017: Auto-renders Mermaid to SVG
├── apps/
│   ├── api/             # Go Backend (Clean Architecture)
│   │   ├── cmd/main.go          # Entry point
│   │   ├── internal/            # Core Business Logic
│   │   │   ├── database/        # Database Access Layer (sqlc)
│   │   │   ├── domain/          # Pure Entities (Independent)
│   │   │   ├── transport/       # Transport Layer (gRPC/Proto)
│   │   │   ├── mappers/         # Model Mappers (Goverter)
│   │   │   ├── adapters/        # Entry/Exit Adapters
│   │   │   ├── usecase/         # Use Case Interactors
│   │   │   └── infra/           # Shared Infrastructure
│   │   ├── migrations/          # Atlas managed SQL migrations (ADR-006)
│   │   ├── atlas.hcl            # Migration configuration
│   │   └── go.mod
│   ├── web/             # Next.js Dashboard
│   └── mobile/          # Expo (React Native)
├── proto/               # THE SOURCE OF TRUTH (ADR-012, 014)
│   ├── dilocash/v1/
│   │   └── api.proto    # Consolidated API Contract
│   └── buf.yaml         # Buf configuration (ADR-040)
├── packages/
│   ├── ui/              # Shared UI components (Tamagui)
│   ├── schema/          # Shared Zod/TS schemas
│   └── database/        # Database Schema & Queries
│       ├── schema.sql           # SQL Schema Source
│       ├── queries.sql          # sqlc Queries
│       └── sqlc.yaml            # sqlc Configuration
├── docs/
│   ├── adr/             # Architecture Decision Records
│   ├── diagrams/        # Manual Diagrams
│   └── generated/       # AUTOMATIC DIAGRAMS (Mermaid/SVG)
├── infra/               # Pulumi IaC (ADR-023, ADR-041)
├── Makefile             # Global task runner
├── buf.gen.yaml         # Protobuf generation config (ADR-040)
├── docker-compose.yaml  # Local infrastructure (Postgres, Redis)
└── turbo.json           # Turborepo configuration


```

## 🔐 Authentication & Identity

Dilocash follows modern security standards to ensure your financial data is protected.

- **Social Auth:** Seamless login via Google, Facebook, and Apple (OAuth 2.0 / OIDC).
- **JWT Validation:** The Go backend operates as a stateless resource server, verifying cryptographically signed tokens for every request.
- **Biometrics:** The Expo mobile app supports **FaceID/Fingerprint** to unlock the app locally.
- **Universal Identity:** Your single Dilocash account is the hub for your Web, Mobile, Telegram, and WhatsApp connectors.

---

## 📜 ADRs (Architecture Decision Records)

We document our "Why." All architectural choices are logged in `docs/adr/`.

- **ADR 001:** [Record Architecture Decisions](docs/adr/0001-record-architecture-decisions.md)
- **ADR 002:** [Monorepo Strategy for Universal UI](docs/adr/0002-monorepo-strategy-for-universal-ui.md)
- **ADR 003:** [Golang Backend with Clean Architecture](docs/adr/0003-golang-backend-with-clean-architecture.md)
- **ADR 004:** [Open Core Extension Model](docs/adr/0004-open-core-extension-model.md)
- **ADR 005:** [Voice-to-Ledger (NLU) Pipeline](docs/adr/0005-voice-to-ledger-nlu-pipeline.md)
- **ADR 006:** [Database Integrity and Monetary Precision](docs/adr/0006-database-integrity-and-monetary-precision.md)
- **ADR 007:** [Multi-Channel Entry Points (Adapters)](docs/adr/0007-multi-channel-entry-points-adapters.md)
- **ADR 008:** [External Identity Mapping](docs/adr/0008-external-identity-mapping.md)
- **ADR 009:** [Ephemeral State Management with Redis](docs/adr/0009-ephemeral-state-management-with-redis.md)
- **ADR 010:** [Identity Management via OIDC](docs/adr/0010-identity-management-via-oidc.md)
- **ADR 011:** [Selecting Supabase as Primary IdP](docs/adr/0011-selecting-supabase-as-primary-idp.md)
- **ADR 012:** [Contract-First API Development](docs/adr/0012-contract-first-api-development.md)
- **ADR 013:** [Diagram-Driven Contract Generation](docs/adr/0013-diagram-driven-contract-generation.md)
- **ADR 014:** [Single Source of Truth for API Contracts](docs/adr/0014-single-source-of-truth-for-api-contracts.md)
- **ADR 015:** [Decoupled Documentation Assets](docs/adr/0015-decoupled-documentation-assets.md)
- **ADR 016:** [Decentralized Diagram Maintenance](docs/adr/0016-decentralized-diagram-maintenance.md)
- **ADR 017:** [Automated Documentation Rendering](docs/adr/0017-automated-documentation-rendering.md)
- **ADR 018:** [Repository Bifurcation Strategy (Open Core)](docs/adr/0018-repository-bifurcation-strategy-open-core.md)
- **ADR 019:** [Brand Asset Protection](docs/adr/0019-brand-asset-protection.md)
- **ADR 020:** [Cost-Optimized Scaling](docs/adr/0020-cost-optimized-scaling.md)
- **ADR 021:** [Dual-Licensing and Repository Separation](docs/adr/0021-dual-licensing-and-repository-separation.md)
- **ADR 022:** [Hybrid Infrastructure Management](docs/adr/0022-hybrid-infrastructure-management.md)
- **ADR 023:** [Infrastructure Language Standardization](docs/adr/0023-infrastructure-language-standardization.md)
- **ADR 024:** [Managed Storage Policy](docs/adr/0024-managed-storage-policy.md)
- **ADR 025:** [Informed Consent Architecture](docs/adr/0025-informed-consent-architecture.md)
- **ADR 026:** [Data Retention and Improvement Policy](docs/adr/0026-data-retention-and-improvement-policy.md)
- **ADR 027:** [Monetization Strategy](docs/adr/0027-revised-monetization-strategy.md)
- **ADR 028:** [Repository Error Handling](docs/adr/0028-repository-error-handling.md)
- **ADR 029:** [Graceful Shutdown Protocol](docs/adr/0029-graceful-shutdown-protocol.md)
- **ADR 030:** [Communication Protocol Selection](docs/adr/0030-communication-protocol-selection.md)
- **ADR 031:** [Repository and Organization Governance](docs/adr/0031-repository-and-organization-governance.md)
- **ADR 032:** [Bootstrapped Infrastructure Cost Management](docs/adr/0032-bootstrapped-infrastructure-cost-management.md)
- **ADR 033:** [Idempotency Strategy for the Financial Ledger](docs/adr/0033-idempotency-strategy-for-the-financial-ledger.md)
- **ADR 034:** [Global i18n & l10n Strategy (Full-Stack)](docs/adr/0034-global-i18n-l10n-strategy-full-stack.md)
- **ADR 035:** [Rate Limiting and Resilience](docs/adr/0035-rate-limiting-and-resilience.md)
- **ADR 036:** [OWASP Top 10 Security and Quality Gate](docs/adr/0036-owasp-top-10-security-and-quality-gate.md)
- **ADR 037:** [Testing Strategy & E2E Frameworks](docs/adr/0037-testing-strategy-e2e-frameworks.md)
- **ADR 038:** [gRPC API Versioning Strategy](docs/adr/0038-grpc-api-versioning-strategy.md)
- **ADR 039:** [AI-Assisted Manual Diagram Maintenance](docs/adr/0039-ai-assisted-manual-diagram-maintenance.md)
- **ADR 040:** [Adoption of Buf for Protobuf Management](docs/adr/0040-adoption-of-buf-for-protobuf-management.md)
- **ADR 041:** [Adoption of Fargate for gRPC Support](docs/adr/0041-adoption-of-fargate-for-grpc-support.md)
- **ADR 042:** [Adoption of Goverter for Automated Model Mapping](docs/adr/0042-adoption-of-goverter-for-mapping.md)

---

## 📄 License

Dilocash (Core) is licensed under the [MIT License](https://www.google.com/search?q=LICENSE).
