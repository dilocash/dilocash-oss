# ADR-005: Voice-to-Ledger (NLU) Pipeline

- **Status**: Accepted
- **Date**: 2026-02-03
- **Authors**: @jalbarran
- **Technical Domain**: AI

## 1. Context and Problem Statement

The primary differentiator for **Dilocash** is the "Talk to Track" feature. Manual data entry is the main friction point for financial tracking, and we need a reliable way to convert natural speech into structured financial records.

## 2. Decision Drivers

- Accuracy in intent extraction.
- Speed of transcription and parsing.
- Prevention of "AI Hallucinations" in financial data.

## 3. Consequences

Implement an asynchronous **Voice-to-Ledger** pipeline using OpenAI **Whisper** and **GPT-4o-mini**.

1. **Whisper**: Transcribe audio to text.
2. **GPT-4o-mini**: Use "Structured Outputs" to map text to the Dilocash schema.
3. **Smart Confirmation**: The UI must show interpreted data for user approval before persistence.

- **Positive:** Frictionless UX for "on-the-go" logging.
- **Negative:** Introduces external API dependency and associated costs.
- **Negative:** Requires robust error handling for failed NLU mapping.

---
