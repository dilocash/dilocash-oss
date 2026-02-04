# ADR-005: Voice-to-Ledger (NLU) Pipeline

- **Status**: Accepted
- **Date**: 2026-02-03
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

The primary differentiator for **Dilocash** is the "Talk to Track" feature. Manual data entry is the main friction point for financial tracking.

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: See bullets below**

Implement an asynchronous **Voice-to-Ledger** pipeline:

1. **Whisper (OpenAI)**: Transcribe audio to text.
2. **GPT-4o-mini**: Use "Structured Outputs" to map text to the Dilocash Transaction schema.
3. **Smart Confirmation**: The UI must show the interpreted data for user approval before persistence to avoid AI hallucinations.

### Technical Implementation Details

[Refer to codebase or diagrams for implementation specifics.]

## 5. Consequences

### Positive (Pros)

- Frictionless UX for "on-the-go" logging.
- Requires robust error handling for failed NLU mapping.

### Negative (Cons/Risks)

[TBD]

## 6. Pros and Cons of Options

### [Option 1]

[TBD]

---
