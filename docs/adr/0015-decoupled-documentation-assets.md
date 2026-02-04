# ADR-015: Decoupled Documentation Assets

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: AI

## 1. Context and Problem Statement

Embedding large Mermaid code blocks directly into the README makes the file difficult to maintain and reduces the visibility of architectural changes in Git diffs.

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: All Mermaid diagrams will reside as independent `.mmd` files in the `docs/diagrams/` directory.**

All Mermaid diagrams will reside as independent `.mmd` files in the `docs/diagrams/` directory.

### Technical Implementation Details

[Refer to codebase or diagrams for implementation specifics.]

## 5. Consequences

### Positive (Pros)

- \* **Pros:** Cleaner README; granular version control for architecture; diagrams can be reused in different documents.
- **Cons:** Requires a link-click to view the diagram on some platforms unless converted to SVG/PNG.

### Negative (Cons/Risks)

[TBD]

## 6. Pros and Cons of Options

### [Option 1]

[TBD]

---
