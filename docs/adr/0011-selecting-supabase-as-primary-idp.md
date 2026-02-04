# ADR-011: Selecting Supabase as Primary IdP

- **Status**: Accepted
- **Context:** We need a scalable, OIDC-compliant provider to handle social logins (Google/Facebook) and secure session management.
- **Decision:** Use **Supabase Auth**.
- **Mechanism:** \* The **Frontend** (Next.js/Expo) uses the Supabase SDK to authenticate.
- The **Go Backend** acts as a stateless "Resource Server" that validates the JWT signature using the Supabase `JWT_SECRET`.
- **Identity Mapping:** The `sub` claim in the JWT is used as the `id` in our `users` table.
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

We need a scalable, OIDC-compliant provider to handle social logins (Google/Facebook) and secure session management.

- **Decision:** Use **Supabase Auth**.
- **Mechanism:** \* The **Frontend** (Next.js/Expo) uses the Supabase SDK to authenticate.
- The **Go Backend** acts as a stateless "Resource Server" that validates the JWT signature using the Supabase `JWT_SECRET`.
- **Identity Mapping:** The `sub` claim in the JWT is used as the `id` in our `users` table.

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: See bullets below**

Use **Supabase Auth**.

- **Mechanism:** \* The **Frontend** (Next.js/Expo) uses the Supabase SDK to authenticate.
- The **Go Backend** acts as a stateless "Resource Server" that validates the JWT signature using the Supabase `JWT_SECRET`.
- **Identity Mapping:** The `sub` claim in the JWT is used as the `id` in our `users` table.

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
