# ADR-009: Ephemeral State Management with Redis

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Infra

## 1. Context and Problem Statement

For multi-channel handshakes (like the 6-digit PIN for Telegram linking), we need a fast, ephemeral storage solution. Storing these short-lived, high-frequency tokens in PostgreSQL would create unnecessary disk I/O and bloat the primary database.

## 2. Decision Drivers

- Performance and low latency for token validation.
- Automatic expiration of temporary data.
- Reduction of load and bloat in the primary relational database.

## 3. Consequences

Use **Redis** as a secondary, ephemeral data store for short-lived state.

- **Positive:** Significant reduction in load on PostgreSQL.
- **Positive:** Native TTL support ensures security tokens Automatically expire after a few minutes.
- **Negative:** Adds a new architectural dependency that must be managed in Docker and production environments.

---
