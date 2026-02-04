# ADR-006: Database Integrity and Monetary Precision

- **Status**: Accepted
- **Date**: 2026-02-03
- **Authors**: @jalbarran
- **Technical Domain**: Infra

## 1. Context and Problem Statement

Financial data cannot tolerate rounding errors. Furthermore, the system must prevent duplicate entries from network retries.

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: See bullets below**

1. **Precision**: Use PostgreSQL `NUMERIC` for all currency amounts. In Go, use `shopspring/decimal`.
2. **Idempotency**: All transaction creation requests must include an `X-Idempotency-Key` header.
3. **Migrations**: Use **Atlas** for declarative database schema management.

### Technical Implementation Details

[Refer to codebase or diagrams for implementation specifics.]

## 5. Consequences

### Positive (Pros)

- 100% mathematical accuracy.

* Safe retries in poor network conditions (common in LatAm mobile use).

### Negative (Cons/Risks)

[TBD]

## 6. Pros and Cons of Options

### [Option 1]

[TBD]

---
