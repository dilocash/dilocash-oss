# ADR-018: Repository Bifurcation Strategy (Open Core)

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

To support a sustainable business model, the project requires a clear separation between the community-driven core and proprietary value-adds.

## 2. Decision Drivers

- Included in Context

## 3. Considered Options

- **Option 1**: Proposed implementation.

## 4. Decision Outcome

**Chosen Option: See bullets below**

Rename the primary repository to `dilocash-oss`.

- **Scope of OSS:** NLU engine, basic adapters (Telegram), and core ledger.
- **Scope of Premium:** Advanced analytics, multi-user household management, and enterprise bank sync (managed in a private repository).
- **Interface Strategy:** Use Go interfaces in `oss` to allow "Pluggable" premium modules.

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
