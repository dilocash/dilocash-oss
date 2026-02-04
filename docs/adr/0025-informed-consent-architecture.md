# ADR-025: Informed Consent Architecture

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Infra

## 1. Context and Problem Statement

Capturing voice data and financial records requires explicit user consent to meet global privacy standards (GDPR/LGPD).

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: See bullets below**

Implement a blocking "Consent Gate" across all entry points.

- **Database:** Add a `terms_accepted_at` timestamp to the `users` table.
- **Logic:** The `IntentService` must reject requests from users where `terms_accepted_at` is NULL.
- **UI:** Display a summary of AI data usage (OpenAI disclosure) directly on the consent screen.

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
