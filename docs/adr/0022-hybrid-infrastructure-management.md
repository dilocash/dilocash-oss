# ADR-022: Hybrid Infrastructure Management

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Infra

## 1. Context and Problem Statement

Local development must be accessible, but production must be reproducible and automated.

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: See bullets below**

- **Development:** Use `docker-compose` as the source of truth for local container orchestration.
- **Production/Staging:** Use **Pulumi (Go SDK)** for Infrastructure as Code.
- **Benefit:** Unified language (Go) for both logic and infrastructure; zero cost for local development; high reliability for cloud deployment.

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
