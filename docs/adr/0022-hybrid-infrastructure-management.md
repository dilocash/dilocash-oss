# ADR-022: Hybrid Infrastructure Management

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Infra

## 1. Context and Problem Statement

Local development must be accessible and cost-free, while production infrastructure must be reproducible, automated, and secure. We need a way to manage both without duplicating infrastructure definitions.

## 2. Decision Drivers

- Developer experience (Zero-cost local dev).
- Production reliability and reproducibility.
- Unified language for both app logic and infrastructure.

## 3. Consequences

Adopt a hybrid approach: **Docker Compose** for local development and **Pulumi (Go SDK)** for cloud infrastructure.

- **Local:** `docker-compose.yaml` provides a replica of the production stack (Postgres/Redis/API).
- **Cloud:** Pulumi scripts define the VPC, RDS, and ECS resources in AWS.

- **Positive:** Zero hosting costs during the local development and testing phase.
- **Positive:** Staff-level reliability in production through Infrastructure as Code (IaC).
- **Negative:** Requires maintaining two sets of configuration files (Compose and Pulumi).

---
