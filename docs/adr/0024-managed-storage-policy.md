# ADR-024: Managed Storage Policy

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Infra

## 1. Context and Problem Statement

Dilocash processes high volumes of audio data and NLU activity logs. Storing these binary and unstructured assets in a relational database like PostgreSQL would be extremely expensive and degrade database performance.

## 2. Decision Drivers

- Minimize storage costs for high-volume logs and audio.
- Ensure high availability and scalability of asset storage.
- Easy integration with CDNs for future web/mobile usage.

## 3. Consequences

Use **Amazon S3** (or S3-compatible storage like MinIO for local dev) for raw audio files and NLU logs.

- **Implementation:** Managed via Pulumi Go SDK with mandatory **Versioning** and **Public Access Blocking** enabled by default.

- **Positive:** Keeps the primary PostgreSQL database lean and performant.
- **Positive:** Drastically lower costs for storing multi-terabyte audio datasets.
- **Negative:** Adds a new external infrastructure dependency and secret management requirement.

---
