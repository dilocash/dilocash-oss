# ADR-015: Decoupled Documentation Assets

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: AI

## 1. Context and Problem Statement

Embedding large Mermaid code blocks directly into the main `README.md` makes the file difficult to maintain, reduces readability, and clutters Git diffs with visual noise.

## 2. Decision Drivers

- Maintain a clean and readable root README.
- Enable granular version control for architecture artifacts.
- Support diagram reuse across multiple technical documents.

## 3. Consequences

Decouple Mermaid diagrams into independent `.mmd` files within a dedicated `docs/diagrams/` directory.

- **Positive:** Purer Git history for architectural changes.
- **Positive:** README remains concise and focused on high-level goals.
- **Negative:** Requires linking to external files, which might not render natively on all git hosting mirrors.

---
