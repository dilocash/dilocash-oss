# ADR-008: External Identity Mapping

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

Dilocash supports multiple entry points (Telegram, Watch). We need a secure way to map platform-specific IDs (like ChatID) to internal User UUIDs without exposing PII.

## 2. Decision Drivers

- Security and privacy (no storage of raw phone numbers).
- One-time setup for external connectors.
- Prevention of unauthorized identity takeovers.

## 3. Consequences

Implement a **Connection Token Handshake** using short-lived PINs.

- Users generate a 6-digit PIN in the main Web/Mobile app.
- The external adapter (Bot/Watch) validates this PIN to link the `platform_id` to the `user_id`.
- Rely exclusively on platform-provided unique IDs (e.g., Telegram ChatID).

- **Positive:** No Plaintext PII as primary identifiers.
- **Positive:** Secure, user-initiated linking process.
- **Negative:** One-time friction for the user during the initial setup of a new device or bot.

---
