# ADR-003: Golang Backend with Clean Architecture

- **Status**: Accepted
- **Date**: 2026-02-03
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

Fintech applications require high reliability, clear audit trails, and strict data integrity. The backend needs to be performant while remaining decoupled from external services (DB, AI APIs).

## 2. Decision Drivers

- High performance and low memory footprint.
- Testability and separation of business logic from infrastructure.
- Strong typing and reliability.

## 3. Consequences

Implement the backend in **Golang** following **Clean Architecture** principles.

- **Entities**: Pure business models.
- **Use Cases**: Domain-specific orchestration logic.
- **Infrastructure**: Implementations for PostgreSQL, OpenAI, etc.
- **SQLC**: Generate type-safe Go code from pure SQL queries.

- **Positive:** Highly testable through mockable interfaces.
- **Positive:** Low memory footprint compared to JVM-based alternatives.
- **Negative:** Requires more boilerplate code initially than "all-in-one" frameworks.

---
