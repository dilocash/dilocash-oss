# ADR-020: Cost-Optimized Scaling

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Infra

## 1. Context and Problem Statement

Financial accessibility is a core tenet of Dilocash. The system must remain affordable to run for the OSS community.

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: See bullets below**

Prioritize "Serverless" and "Pay-as-you-go" infrastructure.

- **Storage:** Use `sqlc` to keep queries efficient and DB size small.
- **AI:** Default to `gpt-4o-mini` for the OSS core to keep costs 90% lower than standard GPT-4.
- **Handshake:** Redis TTLs must be strictly enforced (5 mins) to keep memory usage in the free tier.

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
