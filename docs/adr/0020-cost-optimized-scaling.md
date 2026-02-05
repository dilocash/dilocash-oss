# ADR-020: Cost-Optimized Scaling

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Infra

## 1. Context and Problem Statement

Financial accessibility is a core tenet of the Dilocash project. The system must remain affordable to host for both the maintainers and OSS community members who wish to self-host.

## 2. Decision Drivers

- Minimize monthly recurring infrastructure costs.
- Support for "Serverless" and "Pay-as-you-go" cloud models.
- Efficient usage of expensive resources (like AI tokens).

## 3. Consequences

Prioritize technical choices that minimize operating expenses through efficient resource usage and serverless-first policies.

- **AI:** Default to `gpt-4o-mini` for the OSS core to reduce costs by ~90%.
- **Storage:** Use `sqlc` for high-efficiency queries to keep database tier requirements minimal.
- **Cache:** Enforce strict Redis TTLs (5 mins) to stay within free infrastructure tiers.

- **Positive:** Lower barrier to entry for self-hosters and contributors.
- **Positive:** Sustainable growth profile for the commercial entity.
- **Negative:** `gpt-4o-mini` may be slightly less accurate than larger models for complex, multi-intent voice notes.

---
