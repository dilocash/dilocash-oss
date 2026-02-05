# ADR-018: Repository Bifurcation Strategy (Open Core)

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

To support a sustainable business model while fostering a community, Dilocash requires a clear strategy for separating community-driven "Core" code from proprietary "Premium" value-adds.

## 2. Decision Drivers

- Sustain project development through monetization.
- Foster community contributions through a high-quality OSS core.
- Prevent accidental leakage of proprietary logic into public repositories.

## 3. Consequences

Bifurcate the codebase into `dilocash-oss` (Public) and `dilocash-premium` (Private), using Go interfaces to permit "pluggable" premium modules.

- **OSS Scope:** NLU engine, basic adapters, and core ledger.
- **Premium Scope:** Advanced analytics, managed bank sync, and multi-user household management.

- **Positive:** Enables a sustainable Open Core business model.
- **Positive:** Encourages external contributions to the foundational engine.
- **Negative:** Requires more complex repository management (submodules or separate packages).

---
