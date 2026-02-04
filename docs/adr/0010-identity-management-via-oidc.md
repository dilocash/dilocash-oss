# ADR-010: Identity Management via OIDC

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

Dilocash requires secure authentication via Email and Social Providers (Google, Facebook). Managing raw credentials and OAuth handshakes manually increases security liability.

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: See bullets below**

Outsource Auth to an **Identity Provider (IdP)** using **OpenID Connect (OIDC)**.

- **Backend:** Go will act as a **Resource Server**, validating JWTs from the IdP using a middleware.
- **Frontend:** Next.js and Expo will use the IdP's SDKs for the "Sign-in" flow.
- **Mapping:** The `sub` (subject) claim from the JWT will be used to link the external identity to our internal `users` table.

### Technical Implementation Details

[Refer to codebase or diagrams for implementation specifics.]

## 5. Consequences

### Positive (Pros)

- Documentation and team alignment.

### Negative (Cons/Risks)

[TBD]

## 6. Pros and Cons of Options

### [Option 1]

[TBD]

---
