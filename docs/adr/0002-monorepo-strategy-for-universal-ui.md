# ADR-002: Monorepo Strategy for Universal UI

- **Status**: Accepted
- **Date**: 2026-02-03
- **Authors**: @jalbarran
- **Technical Domain**: Frontend

## 1. Context and Problem Statement

Dilocash requires a Web interface (Next.js) and a Mobile interface (Android/iOS via Expo). Maintaining separate repos leads to duplicated logic for financial calculations, validation, and styling.

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: See bullets below**

Use a **Turborepo** monorepo structure with **Solito** for cross-platform navigation.

- **apps/web-next**: Next.js App Router for SSR.
- **apps/mobile-expo**: React Native for the Android/iOS client.
- **packages/ui**: Shared components using **Tamagui** for "Write Once, Run Everywhere" styling.

### Technical Implementation Details

[Refer to codebase or diagrams for implementation specifics.]

## 5. Consequences

### Positive (Pros)

- Shared business logic (80% code reuse).
- Unified TypeScript types for the entire stack.

### Negative (Cons/Risks)

[TBD]

## 6. Pros and Cons of Options

### [Option 1]

[TBD]

---
