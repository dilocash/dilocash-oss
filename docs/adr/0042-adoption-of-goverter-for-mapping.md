# ADR-042: Adoption of Goverter for Automated Model Mapping

- **Status**: Accepted
- **Date**: 2026-02-06
- **Authors**: @jalbarran
- **Technical Domain**: Backend / Go Development

## 1. Context and Problem Statement

The Dilocash backend architecture relies on generated code from two primary sources:

1. **Protobuf/gRPC**: Service contracts and messages generated via `buf`.
2. **sqlc**: Database models and query methods generated from SQL.

Manually mapping fields between these two generated types (e.g., converting a database `User` struct to a gRPC `User` message) is tedious, error-prone, and violates the DRY (Don't Repeat Yourself) principle. As the schema grows, maintaining these manual mappers becomes a significant source of technical debt and potential bugs.

## 2. Decision Drivers

- **Reduce Boilerplate**: Eliminate hundreds of lines of manual assignment code.
- **Type Safety**: Ensure that mappings are verified at compile-time.
- **Maintainability**: Automatically handle field additions or changes when possible.
- **Performance**: Use a tool that generates efficient Go code instead of relying on slow reflection-based mapping at runtime.

## 3. Consequences

### The Decision

Standardize on **Goverter** as the primary tool for generating type-safe mappers between Protobuf messages and sqlc models.

- **Generation:** Define mapper interfaces with Goverter annotations.
- **Workflow:** Include `goverter` in the `Makefile` to regenerate mappers whenever `.proto` or `.sql` files change.
- **Explicit Mapping:** Use Goverter's custom mapping capabilities for fields that don't match exactly (e.g., date formats or status enums).

### Implications

- **Positive:** Dramatically reduces the amount of "glue code" required in the repository layer.
- **Positive:** High performance as Goverter produces plain Go code without reflection.
- **Positive:** Clearer separation of concerns; mappers are generated artifacts, not handwritten logic.
- **Negative:** Adds another tool dependency to the developer's local environment and CI/CD.
- **Negative:** Requires learning Goverter's specific syntax for complex mapping scenarios.
