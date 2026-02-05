# ADR-039: AI-Assisted Manual Diagram Maintenance

- **Status**: Accepted
- **Date**: 2026-02-05
- **Authors**: @jalbarran
- **Technical Domain**: General

## 1. Context and Problem Statement

Previous decisions (ADR-013 and ADR-014) aimed for fully automated generation of Mermaid diagrams from `.proto` and `.sql` files. However, specialized generator plugins (like `protoc-gen-mermaid` or `sql-to-mermaid`) are often brittle, poorly maintained, or difficult to integrate into standard CI/CD pipelines without significant overhead.

We need high-quality, readable architectural diagrams that stay in sync with the code without the technical debt of managing niche documentation generators.

## 2. Decision Drivers

- Reduce complexity in the `Makefile` and build toolchain.
- Ensure diagrams are visually optimized for human readability (which automated generators often fail to do).
- Leverage modern AI capabilities (e.g., Antigravity, LLMs) to bridge the gap between code and documentation.

## 3. Consequences

### The Decision

Reject the use of automated Mermaid generation tools. Instead, we will use **AI-Assisted Manual Maintenance** for all architectural and schema diagrams.

- **Process:** When `.proto` or `.sql` schemas are updated, the developer will use AI to generate the corresponding Mermaid `.mmd` code and update the files in `docs/diagrams/` manually.
- **Source of Truth:** The code (`.proto`, `.sql`) remains the ultimate source of truth, but the visual documentation is a curated artifact.

### Implications

- **Positive:** Cleaner, more expressive diagrams that follow human-centric layout principles.
- **Positive:** Dramatically simplified `Makefile` and CI/CD environment (removed multiple specific tool dependencies).
- **Negative:** Increased risk of "Documentation Rot" if developers forget to trigger the AI-update after a schema change.
- **Negative:** Manual step required in the documentation workflow.

---
