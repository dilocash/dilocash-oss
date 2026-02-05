# ADR-033: Idempotency Strategy for the Financial Ledger

- **Status**: Accepted
- **Date**: 2026-02-05
- **Authors**: @jalbarran
- **Technical Domain**: General

## 1. Context and Problem Statement

Network retries (common in Telegram bots and mobile apps) must not result in duplicate financial transactions.

## 2. Decision Drivers

Implement a mandatory `Idempotency-Key` (UUID) in gRPC metadata.

## 3. Consequences

- The backend will store keys in Redis with a 24-hour TTL.
- If a duplicate key is detected, the server returns the cached result without re-processing logic.
