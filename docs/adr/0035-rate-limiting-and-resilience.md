# ADR-035: Rate Limiting and Resilience

- **Status**: Accepted
- **Date**: 2026-02-05
- **Authors**: @jalbarran
- **Technical Domain**: Infra

## 1. Context and Problem Statement

As a public-facing API that orchestrates expensive operations (like AI NLU parsing), Dilocash is vulnerable to both malicious DoS attacks and runaway costs from genuine but excessive usage. We need a way to protect the infrastructure and the project's financial sustainability.

## 2. Decision Drivers

- Protection against brute-force and Denial-of-Service attacks.
- Granular control over "Premium" vs "Community" AI usage limits.
- System stability during traffic spikes.

## 3. Consequences

### The Decision

Implement a two-tier **Token Bucket** rate limiting strategy backed by **Redis**.

1. **User-Level Limit:** Enforce quotas on AI-heavy operations based on the `user_id` to prevent unintentional account-level cost spikes.
2. **IP-Level Limit:** Standard infrastructure protection at the edge to block malicious volumetric traffic.

### Implications

- **Positive:** Predictable infrastructure and AI API monthly costs.
- **Positive:** Significantly higher resiliency during coordinated traffic bursts.
- **Negative:** Potential for "false positives" where heavy-duty users might hit limits during intensive legitimate logging sessions.
- **Negative:** Adds complexity to the gRPC interceptor logic.

---
