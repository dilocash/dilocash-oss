# ADR-040: Adoption of Buf for Protobuf Management

- **Status**: Accepted
- **Date**: 2026-02-06
- **Authors**: @jalbarran
- **Technical Domain**: Backend / Infrastructure

## 1. Context and Problem Statement

Standard `protoc` commands are difficult to maintain, lack built-in linting, and make it hard to enforce breaking change policies. As the service mesh grows, maintaining consistency across multiple `.proto` files using raw shell scripts or manual `protoc` invocations becomes error-prone and brittle.

## 2. Decision Drivers

- Need for a unified tool for linting, formatting, and code generation.
- Requirement to enforce gRPC best practices and API stability.
- Desire for a better developer experience and simplified local setup.
- Integration into the CI/CD pipeline for automated checks.

## 3. Consequences

### The Decision

Standardize on **Buf CLI** for all Protobuf operations, including linting, formatting, and generation.

- **Linting:** Use `buf lint` to enforce style and best practices.
- **Formatting:** Use `buf format` to ensure a consistent look across all proto files.
- **Generation:** Use `buf generate` with `buf.gen.yaml` to replace complex `protoc` invocations.
- **Breaking Changes:** Use `buf breaking` to prevent accidental breaking changes in the API.

### Implications

- **Positive:** Superior developer experience with simplified commands.
- **Positive:** Ensures gRPC best practices via the built-in linter.
- **Positive:** Simplifies the CI/CD pipeline for the `dilocash-oss` core.
- **Positive:** Provides a clear path for managing dependencies via Buf modules.
- **Negative:** Requires developers to install the `buf` CLI.
- **Negative:** Initial migration effort to create `buf.yaml` and `buf.gen.yaml` configurations.
