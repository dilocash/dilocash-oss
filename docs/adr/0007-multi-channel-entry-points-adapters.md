# ADR-007: Multi-Channel Entry Points (Adapters)

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

Users need to log expenses where they happen—on their wrist, in a chat app, or via the dashboard.

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: See bullets below**

We will decouple the **NLU Engine** from the HTTP transport layer. We will implement "Input Adapters" for external services.

- **Telegram/WhatsApp:** Use Webhooks to feed the Go backend.
- **WearOS:** Use a lightweight gRPC or REST client from the Android Watch app.
- **Internal Service:** All adapters call a unified `usecase.ProcessIntent(payload)` function.

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
