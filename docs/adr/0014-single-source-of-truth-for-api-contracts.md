# ADR-014: Single Source of Truth for API Contracts

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

We need visual documentation that never diverges from the actual implementation. Manual diagramming is a high-maintenance overhead that often results in stale documentation.

## 2. Decision Drivers

- Maintain 100% accuracy in documentation.
- Reduce manual overhead for developers.
- Visual-first understanding of the system architecture.

## 3. Consequences

Use **Protobuf (.proto)** as the single master definition for both code and diagrams.

- **Doc-Gen:** Generate Mermaid diagrams from `.proto` files using `protoc-gen-mermaid`.
- **Enforcement:** CI/CD will fail if diagrams are not regenerated following contract changes.

- **Positive:** Zero-maintenance diagrams — they update automatically with code.
- **Positive:** High confidence for new developers reading the docs.
- **Negative:** Limited flexibility in the visual layout of automatically generated diagrams.

---
