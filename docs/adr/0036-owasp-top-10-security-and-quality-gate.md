# ADR-036: OWASP Top 10 Security and Quality Gate

- **Status**: Accepted
- **Date**: 2026-02-05
- **Authors**: @jalbarran
- **Technical Domain**: General

## 1. Context and Problem Statement

As a fintech app, security and code quality are non-negotiable.

## 2. Decision Drivers

Implement security best practices to protect against common vulnerabilities.

## 3. Consequences

**Security:** Strict use of parameterized queries and JWT validation in gRPC interceptors.

- **SAST:** Integrate `gosec` and `SonarQube` into GitHub Actions. Merges are blocked if the "Quality Gate" or critical security scans fail.
