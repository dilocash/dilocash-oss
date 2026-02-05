# ADR-011: Selecting Supabase as Primary IdP

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

We need a specific, scalable, and OIDC-compliant provider to handle social logins and secure session management for the Dilocash ecosystem.

## 2. Decision Drivers

- Cost-effective (free tier for the OSS community).
- Ease of integration with Next.js and Go.
- Native support for multiple social providers out-of-the-box.

## 3. Consequences

Select **Supabase Auth** as the primary Identity Provider for Dilocash.

- **Mechanism:** The **Frontend** uses Supabase SDK; the **Go Backend** validates JWT signatures using the shared secret.
- **Mapping:** Use the Supabase `sub` claim for our internal identification.

- **Positive:** Rapid implementation of Google/Facebook login.
- **Positive:** Leverages Supabase's managed infrastructure and security patches.
- **Negative:** Tighter coupling with the Supabase ecosystem for the core authentication layer.

---
