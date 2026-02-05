# ADR-025: Informed Consent Architecture

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Infra

## 1. Context and Problem Statement

Capturing voice data and financial records requires explicit user consent to meet global privacy standards such as GDPR (Europe) and LGPD (Brazil). We need a system-wide mechanism to ensure no sensitive data is processed without a record of user agreement.

## 2. Decision Drivers

- Legal compliance with global privacy regulations.
- User trust and transparency regarding AI data usage.
- Blocking enforcement across all entry points (Web, Mobile, Bots).

## 3. Consequences

Implement a blocking "Consent Gate" enforced at the application layer.

- **Database:** Track `terms_accepted_at` in the `users` table.
- **Enforcement:** The `IntentService` logic must reject requests from users without a valid consent timestamp.
- **UI:** Integrate an explicit OpenAI/AI disclosure during the onboarding "Consent Screen."

- **Positive:** Guarantees legal compliance for voice data handling.
- **Positive:** Builds user confidence through clear data-handling disclosures.
- **Negative:** Adds a mandatory step to the user onboarding flow, slightly increasing friction.

---
