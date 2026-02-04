# ADR-003: Golang Backend with Clean Architecture

- **Status**: Accepted
- **Date**: 2026-02-03
- **Authors**: @jalbarran
- **Technical Domain**: Infra

## 1. Context and Problem Statement

Fintech applications require high reliability, clear audit trails, and strict data integrity. The backend needs to be performant while remaining decoupled from external services (DB, AI APIs).

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: See bullets below**

Implement the backend in **Golang** using **Clean Architecture** (Uncle Bob).

- **Entities**: Business models (Transaction, Account).
- **Use Cases**: Orchestration (ProcessVoiceTransaction, GetReport).
- **Infrastructure**: Implementations of interfaces (PostgreSQL, OpenAI).
- **SQLC**: We will generate type-safe Go code from pure SQL.

### Technical Implementation Details

[Refer to codebase or diagrams for implementation specifics.]

## 5. Consequences

### Positive (Pros)

- Highly testable (mockable interfaces).
- Low memory footprint compared to JVM-based alternatives.

### Negative (Cons/Risks)

[TBD]

## 6. Pros and Cons of Options

### [Option 1]

[TBD]

---
