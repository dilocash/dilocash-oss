# ADR-030: Communication Protocol Selection

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

Choosing a protocol for client-server communication that balances performance, strictness, and future-proofing for mobile and AI features. We need a communication protocol that ensures type safety, high performance for mobile users, and supports future streaming for voice processing.

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: Standardize on **gRPC** using **Protocol Buffers (proto3)**.**

Standardize on **gRPC** using **Protocol Buffers (proto3)**.

### Technical Implementation Details

[Refer to codebase or diagrams for implementation specifics.]

## 5. Consequences

### Positive (Pros)

- **Positive:** Strict contracts between frontend/backend; smaller binary payloads; native streaming.
- **Negative:** Requires gRPC-compatible clients; slightly higher initial setup complexity than REST.

### Negative (Cons/Risks)

[TBD]

## 6. Pros and Cons of Options

### [Option 1]

[TBD]

---
