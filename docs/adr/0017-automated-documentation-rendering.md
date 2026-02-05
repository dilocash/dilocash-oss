# ADR-017: Automated Documentation Rendering

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: AI

## 1. Context and Problem Statement

Manual documentation updates are often neglected, leading to "Documentation Rot" where architectural diagrams diverge from the actual code. We need a frictionless way to maintain visual architectural fidelity.

## 2. Decision Drivers

- Maintain 100% accuracy in visual documentation.
- Zero manual effort for developers during diagram updates.
- Support for version-controlled, text-based diagram sources.

## 3. Consequences

Implement a **GitHub Action** to automatically render Mermaid `.mmd` files into `.svg` images upon every push to the `main` branch.

- **Positive:** README and documentation always reflect the current state of the architecture.
- **Positive:** Zero manual overhead for developers; diagrams are updated by simply editing text.
- **Negative:** Slightly longer CI times (usually < 30s).
- **Negative:** Repository size grows slightly over time due to storage of image assets.

---
