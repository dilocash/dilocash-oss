# ADR-034: Global i18n & l10n Strategy (Full-Stack)

- **Status**: Accepted
- **Date**: 2026-02-05
- **Authors**: @jalbarran
- **Technical Domain**: General

## 1. Context and Problem Statement

Dilocash serves a global audience. Currency, dates, and labels must adapt to the user’s locale across Go, React, and React Native.

## 2. Decision Drivers

Currency, dates, and labels must adapt to the user’s locale across Go, React, and React Native.

## 3. Consequences

**Backend:** Use `golang.org/x/text` and ISO 4217 for currencies.

**Frontend (Web/Mobile):** Standardize on **`i18next`** and the native **`Intl`** JS API for formatting.

**Sync:** Use the `Accept-Language` gRPC header to synchronize the UI language with server-generated messages.
