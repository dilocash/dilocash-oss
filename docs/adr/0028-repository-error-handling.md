# ADR-028: Repository Error Handling

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

Database-level errors such as connection timeouts or network loss can degrade the API's responsiveness or cause partial state updates if not handled gracefully. We need a standardized way to manage these failures.

## 2. Decision Drivers

- Resilience against intermittent infrastructure failures.
- Prevent resource leaks (orphaned database connections).
- Ensure the API returns clear results during high-load periods.

## 3. Consequences

Standardize on context-aware repository methods and explicit error wrapping.

- **Context:** All repository methods must accept a `context.Context` to support cancellation and timeouts.
- **Responsibility:** The repository layer must wrap database errors to provide domain-level context to the use-cases.

- **Positive:** Prevents "hanging" requests from consuming backend resources.
- **Positive:** Enables much easier debugging of database-related failures via structured logs.
- **Negative:** Adds a small amount of boilerplate to every database query method.

---
