# ADR-035: Rate Limiting and Resilience

- **Status**: Accepted
- **Date**: 2026-02-05
- **Authors**: @jalbarran
- **Technical Domain**: General

## 1. Context and Problem Statement

Protection against DoS attacks and managing OpenAI API costs is critical.

## 2. Decision Drivers

Implement a two-tier **Token Bucket** rate limiter using Redis:

## 3. Consequences

1. **User Level:** Limit requests per `user_id` (AI usage control).
2. **IP Level:** Global infrastructure protection.
