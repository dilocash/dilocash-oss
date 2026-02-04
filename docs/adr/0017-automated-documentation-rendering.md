# ADR-017: Automated Documentation Rendering

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: AI

## 1. Context and Problem Statement

Manual documentation updates are often neglected, leading to "Doc Rot." We need a friction-less way to maintain visual architectural fidelity.

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: Implement a **GitHub Action** to render Mermaid `.mmd` files into `.svg` images.**

Implement a **GitHub Action** to render Mermaid `.mmd` files into `.svg` images.

### Technical Implementation Details

[Refer to codebase or diagrams for implementation specifics.]

## 5. Consequences

### Positive (Pros)

- **Pros:** README always reflects the current `main` branch state; Zero manual overhead for developers.
- **Cons:** Slightly longer CI times (usually < 30s); repository size grows slightly due to binary/image assets.

### Negative (Cons/Risks)

[TBD]

## 6. Pros and Cons of Options

### [Option 1]

[TBD]

---
