# ADR-007: Multi-Channel Entry Points (Adapters)

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

Users need to log expenses where they happen—on their wrist, in a chat app, or via the dashboard. We need a way to support multiple "Input Adapters" without duplicating parsing logic.

## 2. Decision Drivers

- Decouple transport layers from business logic.
- Support for Telegram, WhatsApp, WearOS, and Web.
- Unified entry point for the NLU engine.

## 3. Consequences

Decouple the **NLU Engine** from the transport layer using the **Adapter Pattern**.

- All adapters (Bot, Watch, Web) funnel into a unified `usecase.ProcessIntent(payload)` function.
- **Telegram/WhatsApp:** Use Webhooks as intake.
- **WearOS:** Use a lightweight gRPC client.

- **Positive:** Adding a new entry point (e.g., Slack) requires zero changes to core logic.
- **Positive:** Consistent parsing behavior across all platforms.
- **Negative:** Requires careful versioning of the internal Intent payload to avoid breaking adapters.

---
