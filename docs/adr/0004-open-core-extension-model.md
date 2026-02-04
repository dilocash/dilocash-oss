# ADR-004: Open Core Extension Model

- **Status**: Accepted
- **Date**: 2026-02-03
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

We want to provide an Open Source core while monetizing premium features (AI NLU, Bank Sync). We need to separate proprietary code without breaking the OSS build.

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: See bullets below**

Adopt an **Open Core** model using **Git Submodules**.

- Public Repo (`dilocash-oss`) contains the core engine.
- Private Repo (`dilocash-premium`) is linked as a submodule in the `/premium` directory.
- Use Go **Dependency Injection** to swap OSS implementations with Premium ones at runtime based on environment flags.

### Technical Implementation Details

[Refer to codebase or diagrams for implementation specifics.]

## 5. Consequences

### Positive (Pros)

- IP Protection for premium logic.

* Increased complexity in CI/CD pipelines.

### Negative (Cons/Risks)

[TBD]

## 6. Pros and Cons of Options

### [Option 1]

[TBD]

---
