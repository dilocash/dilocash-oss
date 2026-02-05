# ADR-006: Database Integrity and Monetary Precision

- **Status**: Accepted
- **Date**: 2026-02-03
- **Authors**: @jalbarran
- **Technical Domain**: Infra

## 1. Context and Problem Statement

Financial data cannot tolerate rounding errors or duplicates from network retries. We need a robust storage strategy to ensure 100% accuracy and idempotency.

## 2. Decision Drivers

- Zero tolerance for rounding errors.
- System reliability in poor network conditions (e.g., LatAm mobile).
- Reproducible and version-controlled schema management.

## 3. Consequences

Use PostgreSQL `NUMERIC` types, gRPC Idempotency keys, and **Atlas** for migrations.

- **Precision**: PostgreSQL `NUMERIC` + Go `shopspring/decimal`.
- **Idempotency**: Require `X-Idempotency-Key` for all transaction creation.
- **Migrations**: Declarative schema management via **Atlas**.

- **Positive:** 100% mathematical accuracy.
- **Positive:** Safe retries in poor network conditions without double-charging.
- **Negative:** Slightly higher development overhead for handling decimal types compared to floats.

---
