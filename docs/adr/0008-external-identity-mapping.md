# ADR-008: External Identity Mapping

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

Dilocash supports multiple entry points. We need a secure way to map platform-specific IDs to internal User UUIDs.

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: See bullets below**

We will use a **Connection Token Handshake**.

- Users generate a 6-digit PIN in the authenticated web/mobile app.
- The external adapter (Bot/Watch) validates this PIN to create a permanent mapping in the `connections` table.
- Plaintext phone numbers or PII (Personally Identifiable Information) will not be used as primary identifiers; we rely on the unique IDs provided by the platform APIs (e.g., Telegram ChatID).

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
