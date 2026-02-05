# ADR-032: Bootstrapped Infrastructure Cost Management

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: General

## 1. Context and Problem Statement

During the bootstrapping phase, we need to minimize fixed operational costs without compromising on professional standards or development velocity.

## 2. Decision Drivers

- Sustain project development with $0 fixed costs for core tools.
- Maintain professional CI/CD and repository orchestration.
- Allow for easy scaling once funding or revenue is established.

## 3. Consequences

Standardize on the **GitHub Free Organization** tier for project management and hosting.

- **Positive:** Provides the necessary private/public repo split at zero cost.
- **Positive:** Unlimited GitHub Actions minutes for public repositories.
- **Negative:** Limited private repo storage and Actions minutes compared to paid tiers.

---
