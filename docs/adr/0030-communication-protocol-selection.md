# ADR-030: Communication Protocol Selection

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

We need a communication protocol for client-server interaction that balances performance, strictness, and future-readiness for streaming-heavy AI features.

## 2. Decision Drivers

- Type safety across Go, TypeScript, and Dart (Mobile).
- High performance for mobile users on varied network qualities.
- Native support for data streaming (required for real-time voice processing).

## 3. Consequences

Standardize on **gRPC** using **Protocol Buffers (proto3)** as the primary communication layer.

- **Positive:** Strict, enforceable contracts between frontends and the backend.
- **Positive:** Native support for bi-directional streaming for advanced AI features.
- **Positive:** Significant reduction in payload size (~30-50% compared to JSON).
- **Negative:** Requires more complex initial setup (Buf, code generation) than traditional REST.

---
