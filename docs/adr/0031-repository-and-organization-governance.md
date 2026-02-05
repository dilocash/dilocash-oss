# ADR-031: Repository and Organization Governance

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: General

## 1. Context and Problem Statement

To manage the Open Core model effectively, we need a governance structure that protects the commercial identity while enabling public community contributions.

## 2. Decision Drivers

- Clear legal boundary between MIT and Proprietary code.
- Minimize "Blast Radius" of potential security breaches.
- Professional organization structure for external stakeholders.

## 3. Consequences

Establish a **GitHub Organization** with a strict Public/Private repository split.

- **Identity:** Transition from a personal account to the `@dilocash-oss` organization.
- **Isolation:** Infrastructure and Premium repositories remain private; the core engine is public.

- **Positive:** Clear separation of concerns for security and licensing.
- **Positive:** Easier management of team permissions and secrets across different repos.
- **Negative:** Adds a small administrative overhead for repository and member management.

---
