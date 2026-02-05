# ADR-036: OWASP Top 10 Security and Quality Gate

- **Status**: Accepted
- **Date**: 2026-02-05
- **Authors**: @jalbarran
- **Technical Domain**: General

## 1. Context and Problem Statement

For a fintech application, security is not an optional feature but a foundational requirement. We must protect user financial data from common vulnerabilities and ensure that every code change maintains high engineering standards.

## 2. Decision Drivers

- Zero-tolerance for SQL Injection or unauthorized data access.
- Compliance with secure coding practices (OWASP Top 10).
- Prevention of common Go-specific security pitfalls.

## 3. Consequences

### The Decision

Force every code contribution through a multi-layered Security and Quality Gate.

- **Coding Standard:** Strict use of parameterized queries via `sqlc` and mandatory JWT verification in all non-public gRPC interceptors.
- **Automation (SAST):** Integrate `gosec`, `SonarQube`, and dependency vulnerability scanning (Snyk/GitHub) into the PR pipeline.
- **Gating:** Merges to `main` are automatically blocked if the "Quality Gate" or critical security scans find high-risk issues.

### Implications

- **Positive:** High architectural confidence and lowered risk of security breaches.
- **Positive:** Automated detection of technical debt and "code smells."
- **Negative:** Slightly longer CI feedback loops for developers during the PR process.
- **Negative:** Initial setup effort for maintaining the SonarQube/SAST scanners.

---
