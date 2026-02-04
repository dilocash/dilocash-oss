# ADR-026: Data Retention and Improvement Policy

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Infra

## 1. Context and Problem Statement

Improving NLU accuracy requires real-world data, but storing raw audio indefinitely creates a privacy risk.

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: See bullets below**

- **Retention:** Raw audio files used for improvement will be deleted after **90 days**.
- **Storage:** Audio and Transcripts will be stored in a separate, "de-identified" S3 bucket.
- **Access:** Only "Contributor" level internal roles can access the scrubbed training sets.

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
