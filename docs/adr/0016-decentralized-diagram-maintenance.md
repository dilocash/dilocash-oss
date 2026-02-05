# ADR-016: Decentralized Diagram Maintenance

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Infra

## 1. Context and Problem Statement

As a monorepo grows, centralized documentation becomes a bottleneck. We need a way to maintain diagrams close to the code they describe while keeping the root documentation clean.

## 2. Decision Drivers

- Granular documentation ownership.
- Cleaner root README.
- Ease of reviewing architectural changes in localized PRs.

## 3. Consequences

Adopt a **Decentralized Embedding** approach where diagrams live near their sub-projects but are linked from a central directory.

- **Positive:** Decoupled documentation reduces merge conflicts on the main README.
- **Positive:** Encourages developers to update relevant docs during feature work.
- **Negative:** Slightly harder to get a "one-page" complete system overview.

---
