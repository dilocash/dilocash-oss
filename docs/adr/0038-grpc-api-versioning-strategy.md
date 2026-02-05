# ADR-038: gRPC API Versioning Strategy

- **Status**: Accepted
- **Date**: 2026-02-05
- **Authors**: @jalbarran
- **Technical Domain**: General

## 1. Context and Problem Statement

We must evolve the API without breaking active mobile or bot clients.

## 2. Decision Drivers

Implement a versioning strategy that ensures the ledger is 100% accurate and the UI is functional.

## 3. Consequences

Use **Package Versioning** (e.g., `package dilocash.transactions.v1`).

- Follow a "No Breaking Changes" policy for minor updates.
- For disruptive changes, introduce a `v2` package while maintaining `v1` side-by-side during a deprecation window.
