# ADR-038: gRPC API Versioning Strategy

- **Status**: Accepted
- **Date**: 2026-02-05
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

As Dilocash evolves, the API contracts will inevitably change. We need a way to introduce new features and structural changes without breaking existing mobile clients or third-party bot integrations that may not update immediately.

## 2. Decision Drivers

- Stability for long-lived mobile and bot clients.
- Clear communication of breaking changes to the ecosystem.
- Minimal overhead for maintaining multiple API versions.

## 3. Consequences

### The Decision

Adopt a strict **Package-Based Versioning** strategy within the Protobuf definitions.

- **Standard:** Use versioned packages (e.g., `package dilocash.transactions.v1`).
- **Policy:** Minor updates must follow a "No Breaking Changes" rule (only additive changes allowed).
- **Evolution:** For disruptive changes, a `v2` package is introduced. Both `v1` and `v2` are supported side-by-side during a defined deprecation window.

### Implications

- **Positive:** Zero downtime for legacy clients during backend upgrades.
- **Positive:** Explicit self-documentation of the API's evolution.
- **Negative:** Increased maintenance effort in the backend to support multiple service implementations.
- **Negative:** Slightly increased binary size for client-side generated code.

---
