# ADR-002: Monorepo Strategy for Universal UI

- **Status**: Accepted
- **Date**: 2026-02-03
- **Authors**: @jalbarran
- **Technical Domain**: Frontend

## 1. Context and Problem Statement

Dilocash requires a Web interface (Next.js) and a Mobile interface (Android/iOS via Expo). Maintaining separate repositories leads to duplicated logic for financial calculations, validation, and styling.

## 2. Decision Drivers

- Maximize code reuse between Web and Mobile.
- Ensure type safety across the entire stack.
- Simplify dependency management and deployment orchestration.

## 3. Consequences

Use a **Turborepo** monorepo structure with **Solito** for cross-platform navigation and **Tamagui** for shared UI.

- **apps/web-next**: Next.js App Router for SSR.
- **apps/mobile-expo**: React Native for the Android/iOS client.
- **packages/ui**: Shared components using "Write Once, Run Everywhere" styling.

- **Positive:** Shared business logic results in ~80% code reuse.
- **Positive:** Unified TypeScript types across the entire stack.
- **Negative:** Increased complexity in the initial development environment setup.

---
