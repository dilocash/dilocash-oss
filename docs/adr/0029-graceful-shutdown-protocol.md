# ADR-029: Graceful Shutdown Protocol

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

Abruptly terminating the backend process during a deployment or restart can leave database transactions in a partial state or drop active client connections mid-request, leading to data inconsistency.

## 2. Decision Drivers

- Data integrity during system updates.
- Zero-downtime deployment compatibility.
- Cleaner resource cleanup (closing DB connections, flushing logs).

## 3. Consequences

Implement a formalized Graceful Shutdown listener for `SIGINT` and `SIGTERM` signals.

1. **Stop Listening:** The server immediately stops accepting new requests.
2. **Drain:** Wait for active requests to complete or timeout.
3. **Cleanup:** Close the database connection and internal services.

- **Positive:** Ensures no transaction is left "half-finished" during a server restart.
- **Positive:** Smoother experience for users during CI/CD rolling updates.
- **Negative:** Shutdown time is slightly increased to allow for connection draining.

---
