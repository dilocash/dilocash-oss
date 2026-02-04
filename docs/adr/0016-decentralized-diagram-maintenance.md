# ADR-016: Decentralized Diagram Maintenance

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Infra

## 1. Context and Problem Statement

Monorepos can become cluttered if documentation is centralized in a single README. Decoupling diagrams allows for granular PR reviews on architectural changes.

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: See bullets below**

Adopt a **Decentralized Embedding** approach.

- **Storage:** All diagrams live in `docs/diagrams/`.
- **Format:** Use `.md` extensions for diagram files to leverage GitHub's native rendering on a per-file basis.
- **Linking:** The root `README.md` acts as a directory using relative links.

### Technical Implementation Details

[Refer to codebase or diagrams for implementation specifics.]

## 5. Consequences

### Positive (Pros)

- Documentation and team alignment.

### Negative (Cons/Risks)

[TBD]

## 6. Pros and Cons of Options

### [Option 1]

[TBD]

---
