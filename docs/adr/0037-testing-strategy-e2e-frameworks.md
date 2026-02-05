# ADR-037: Testing Strategy & E2E Frameworks

- **Status**: Accepted
- **Date**: 2026-02-05
- **Authors**: @jalbarran
- **Technical Domain**: General

## 1. Context and Problem Statement

We need to ensure the ledger is 100% accurate and the UI is functional.

## 2. Decision Drivers

Implement a testing strategy that ensures the ledger is 100% accurate and the UI is functional.

## 3. Consequences

1. **Backend:** 80% coverage with Unit Tests and Integration Tests (using **Testcontainers-go**).
2. **Web E2E:** Use **Playwright** for automated browser testing.
3. **React Native E2E:** Use **Detox** for "gray-box" testing on mobile emulators.
