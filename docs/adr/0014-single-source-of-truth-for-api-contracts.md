# ADR-014: Single Source of Truth for API Contracts

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

We need visual documentation that never diverges from the actual implementation. Manual diagramming is a high-maintenance overhead.

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: See bullets below**

We will use **Protobuf (.proto)** as the master definition.

- **Code-Gen:** Go interfaces and TS types are derived from `.proto`.
- **Doc-Gen:** Mermaid diagrams are derived from `.proto` via `protoc-gen-mermaid`.
- **Enforcement:** The CI/CD pipeline will fail if the documentation diagrams are not up-to-date with the latest contract changes.

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
