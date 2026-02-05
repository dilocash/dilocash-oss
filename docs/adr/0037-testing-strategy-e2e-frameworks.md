# ADR-037: Testing Strategy & E2E Frameworks

- **Status**: Accepted
- **Date**: 2026-02-05
- **Authors**: @jalbarran
- **Technical Domain**: General

## 1. Context and Problem Statement

To maintain the high reliability expected of a financial ledger, we need a testing strategy that covers everything from low-level logic (unit) to high-level user workflows (E2E) across multiple platforms.

## 2. Decision Drivers

- 100% confidence in ledger mathematical accuracy.
- Prevention of UI regressions during multi-platform releases.
- Speed and stability of the testing pipeline.

## 3. Consequences

### The Decision

Implement a layered testing architecture covering Backend, Web, and Mobile.

1. **Backend:** Minimum 80% coverage using Go `testing` for units and **Testcontainers-go** for full database integration tests.
2. **Web E2E:** Use **Playwright** for high-speed, automated browser testing.
3. **Mobile E2E:** Use **Detox** for "gray-box" automated testing on real emulators/simulators.

### Implications

- **Positive:** Drastically reduced risk of critical production bugs in the ledger or auth flows.
- **Positive:** Faster development cycles as manual QA becomes less necessary.
- **Negative:** Increased CI costs and execution time for full E2E suites.
- **Negative:** Flakiness in mobile E2E tests (Detox) may require ongoing maintenance.

---
