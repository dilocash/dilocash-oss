# ADR-034: Global i18n & l10n Strategy (Full-Stack)

- **Status**: Accepted
- **Date**: 2026-02-05
- **Authors**: @jalbarran
- **Technical Domain**: General

## 1. Context and Problem Statement

Dilocash serves a global audience. Differences in currency symbols, date formats, and localized labels can lead to confusion if not handled correctly. We need a unified strategy to adapt the interface and server responses to the user's locale across the entire stack.

## 2. Decision Drivers

- Consistent localized experience across Web, Mobile, and Bots.
- Support for ISO 4217 currency standards for financial precision.
- "Auto-localization" based on browser/system settings and gRPC headers.

## 3. Consequences

### The Decision

Standardize on industry-standard localization libraries with a shared "Language Sync" protocol.

- **Backend (Go):** Use `golang.org/x/text` for message translation and ISO 4217 for precise currency formatting.
- **Frontend (Web/Mobile):** Standardize on **`i18next`** and the native **`Intl`** JS API.
- **Protocol:** Use the `Accept-Language` gRPC header to synchronize the UI language with server-side generated response messages.

### Implications

- **Positive:** Improved accessibility and user trust for non-English users.
- **Positive:** High precision in financial displays tailored to native user expectations.
- **Negative:** Increased development overhead for maintaining translation bundles (`en.json`, `es.json`, etc.).
- **Negative:** Slightly increased latency for server-side formatting of heavy report payloads.

---
