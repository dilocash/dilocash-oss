# ADR-013: Diagram-Driven Contract Generation

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

We want to minimize the gap between architectural diagrams and implementation. Manual synchronization between Mermaid and Protobuf is prone to human error.

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: See bullets below**

We will adopt a **Diagram-as-Code (DaC)** approach.

- **Master Source:** Mermaid Class Diagrams in the `/docs` folder.
- **Automation:** A custom script (e.g., `scripts/mermaid_to_proto.py`) will parse these diagrams during the `make generate` phase to create the `.proto` definitions.
- **Validation:** If the Mermaid diagram doesn't follow the "Service" naming convention, the build will fail.

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
