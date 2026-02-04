# ADR-009: Ephemeral State Management with Redis

- **Status**: Accepted
- **Context:** Identity linking and rate limiting require high-speed, temporary storage.
- **Decision:** Introduce **Redis** as the "Sidecar" store.
- **Handshake:** Generate 6-digit PINs stored in Redis with a 5-minute TTL.
- **Rate Limiting:** Track AI usage per user to prevent API cost spikes.

- **Consequences:** Reduces load on PostgreSQL and ensures short-lived tokens automatically expire.
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Infra

## 1. Context and Problem Statement

Identity linking and rate limiting require high-speed, temporary storage.

- **Decision:** Introduce **Redis** as the "Sidecar" store.
- **Handshake:** Generate 6-digit PINs stored in Redis with a 5-minute TTL.
- **Rate Limiting:** Track AI usage per user to prevent API cost spikes.

- **Consequences:** Reduces load on PostgreSQL and ensures short-lived tokens automatically expire.

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: See bullets below**

Introduce **Redis** as the "Sidecar" store.

- **Handshake:** Generate 6-digit PINs stored in Redis with a 5-minute TTL.
- **Rate Limiting:** Track AI usage per user to prevent API cost spikes.

- **Consequences:** Reduces load on PostgreSQL and ensures short-lived tokens automatically expire.

### Technical Implementation Details

[Refer to codebase or diagrams for implementation specifics.]

## 5. Consequences

### Positive (Pros)

- Reduces load on PostgreSQL and ensures short-lived tokens automatically expire.

### Negative (Cons/Risks)

[TBD]

## 6. Pros and Cons of Options

### [Option 1]

[TBD]

---
