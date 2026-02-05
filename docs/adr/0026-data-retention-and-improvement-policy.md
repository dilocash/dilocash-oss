# ADR-026: Data Retention and Improvement Policy

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Infra

## 1. Context and Problem Statement

Improving Natural Language Understanding (NLU) accuracy requires access to real-world voice data for training and fine-tuning. However, storing raw financial audio indefinitely creates a significant privacy and security risk.

## 2. Decision Drivers

- Balanced approach between AI improvement and user privacy.
- Minimize the impact of potential data breaches by reducing data life.
- Regulatory compliance regarding data minimization.

## 3. Consequences

Establish a strict 90-day retention and de-identification policy for audio assets.

- **Retention:** Delete original audio files after **90 days**.
- **Isolation:** Store transcripts and audio in a separate, de-identified S3 bucket.
- **Access:** Limit access to scrubbed training sets to "Contributor" roles only.

- **Positive:** Drastically reduces the privacy risk and legal liability of long-term voice storage.
- **Positive:** Enables model improvement without permanent data hoarding.
- **Negative:** Older data is lost for future longitudinal AI training sessions.

---
