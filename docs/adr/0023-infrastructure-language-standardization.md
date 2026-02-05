# ADR-023: Infrastructure Language Standardization

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Infra

## 1. Context and Problem Statement

Switching between different languages for backend logic (Go) and infrastructure (HCL/Terraform) creates cognitive load and prevents the sharing of type definitions and constants.

## 2. Decision Drivers

- Reduce cognitive load for full-stack developers.
- Share domain constants (like database names) directly with infrastructure code.
- Leverage Go's strong type safety for cloud resource definitions.

## 3. Consequences

Standardize on **Go** for all Infrastructure as Code (IaC) using the **Pulumi Go SDK**.

- **Positive:** Catch infrastructure configuration errors at compile-time instead of runtime.
- **Positive:** Shared constants can be imported directly from the app's `internal/domain` layer into the `infra/` layer.
- **Negative:** Developers must learn the Pulumi SDK specific to Go.

---
