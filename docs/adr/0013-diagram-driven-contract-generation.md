# ADR-013: Diagram-Driven Contract Generation

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

We want to minimize the gap between architectural diagrams and implementation. Manual synchronization between Mermaid diagrams and Protobuf definitions is prone to human error and "documentation rot."

## 2. Decision Drivers

- Synchronization between architecture visuals and code.
- "Documentation as Code" (DaC) philosophy.
- Automation of repetitive boilerplate tasks.

## 3. Consequences

Adopt a **Diagram-Driven** approach where Mermaid Class Diagrams can generate or validate Protobuf contracts.

- **Automation:** Use scripts to parse Mermaid definitions during the `make generate` phase.
- **Enforcement:** The build will fail if documentation diagrams diverge from the service naming conventions.

- **Positive:** Architectural diagrams are always an accurate reflection of the code.
- **Negative:** Adds complexity to the build toolchain.
- **Negative:** Requires strict adherence to specific Mermaid syntax to enable parsing.

---
