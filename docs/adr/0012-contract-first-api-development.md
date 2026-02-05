# ADR-012: Contract-First API Development

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

We need a way to ensure that the Go backend and the multi-platform frontends (Web, Mobile, Telegram) stay in sync regarding API method signatures. Ad-hoc JSON changes often lead to runtime errors on disparate platforms.

## 2. Decision Drivers

- Prevent breaking changes between backend and frontend.
- Enable automatic code generation for multiple languages.
- Minimize documentation lag by making the contract the documentation.

## 3. Consequences

Use **Protocol Buffers (Protobuf)** and **Buf** as the source of truth for all API contracts.

- **Code Gen:** Use `buf generate` to produce Go interfaces and TypeScript definitions.
- **Process:** Contracts must be updated and generated before implementation begins.

- **Positive:** Cross-language type safety is guaranteed at compile time.
- **Positive:** Reduced payload size due to binary serialization.
- **Negative:** Slightly higher initial setup and learning curve for contributors unfamiliar with Protobuf.

---
