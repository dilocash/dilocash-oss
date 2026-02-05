# ADR-010: Identity Management via OIDC

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

Dilocash requires secure authentication via Email and Social Providers (Google, Facebook). Managing raw credentials, encryption, and OAuth handshakes manually increases security liability and maintenance overhead.

## 2. Decision Drivers

- Reduce security risk associated with credential storage.
- Speed up development by using industry-standard protocols.
- Provide a seamless social login experience for users.

## 3. Consequences

Outsource identity management to an **Identity Provider (IdP)** using **OpenID Connect (OIDC)**.

- **Backend:** Act as a **Resource Server**, validating incoming JWTs from the IdP.
- **Frontend:** Use IdP-specific SDKs to handle the login flow.
- **Mapping:** Link the external `sub` (subject) claim to our internal `user_id`.

- **Positive:** Zero storage of user passwords in the Dilocash database.
- **Positive:** Compliant with global security standards (SOC2/GDPR) regarding identity.
- **Negative:** Dependency on external IdP uptime and potentially increased latency for token verification.

---
