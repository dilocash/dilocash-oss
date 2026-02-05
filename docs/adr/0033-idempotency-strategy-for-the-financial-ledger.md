# ADR-033: Idempotency Strategy for the Financial Ledger

- **Status**: Accepted
- **Date**: 2026-02-05
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

Network retries are a common occurrence in mobile applications and third-party chat platforms like Telegram or WhatsApp. Without a robust idempotency strategy, these retries can result in duplicate financial transactions, leading to severe data inconsistency and user distrust.

## 2. Decision Drivers

- Guarantee 100% financial transaction integrity.
- Minimize user friction during poor network conditions.
- Prevent race conditions in high-concurrency ledger updates.

## 3. Consequences

### The Decision

Implement a mandatory `Idempotency-Key` (UUID) requirement for all mutation-heavy gRPC calls.

- **Storage:** The backend will store processed keys in **Redis** with a sliding 24-hour expiration (TTL).
- **Mechanism:** If an incoming request contains a key that has already been successfully processed, the server will return the cached result immediately without re-executing any business logic.

### Implications

- **Positive:** Eliminates duplicate transactions caused by platform retries or user double-tapping.
- **Positive:** Improved UX as users receive immediate confirmation for retried requests.
- **Negative:** Adds a dependency on Redis for request lifecycle management.
- **Negative:** Requires clients to strictly generate and manage UUIDs for every mutation request.

---
