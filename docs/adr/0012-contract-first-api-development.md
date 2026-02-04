# ADR-012: Contract-First API Development

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

We need a way to ensure that the Go backend and the multi-platform frontends (Web, Mobile, Telegram) stay in sync regarding API method signatures.

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: See bullets below**

Use **Protocol Buffers (Protobuf)** and **Buf** as the source of truth for all API and Internal Service interfaces.

- **Diagrams:** Mermaid.js will be used in documentation to visualize these contracts.
- **Code Gen:** `buf generate` will be integrated into the Makefile to produce Go interfaces and TypeScript definitions.

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
