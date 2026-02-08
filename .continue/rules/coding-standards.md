---
name: Dilocash Coding Standards
globs: ["**/*.go", "**/*.proto"]
---

# Go & Protobuf Standards

Follow these project-specific standards for all Go development:

## 1. High-Performance Streaming (ADR-042, ADR-043, ADR-044)
- **Never** read entire files (audio/large blobs) into memory.
- Use `io.TeeReader` and `io.Pipe` for concurrent "Fan-Out" (e.g., streaming to S3 and Whisper simultaneously).
- Favor the **Pipe Pattern** to stream data directly from gRPC requests to AWS S3.

## 2. Security & Idempotency
- Every write operation in the ledger must check for an `Idempotency-Key` in gRPC metadata.
- Use parameterized queries for all database interactions (Postgres/pgx).

## 3. Versioning
- Follow **ADR-038**: Never make breaking changes to `v1` protos. Increment the package version to `v2` if a change is disruptive.