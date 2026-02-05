# ADR-021: Dual-Licensing and Repository Separation

- **Status**: Accepted
- **Date**: 2026-02-04
- **Authors**: @jalbarran
- **Technical Domain**: General

## 1. Context and Problem Statement

To build a sustainable ecosystem, we need a legal framework that allows for both free community use and proprietary commercial extensions. Clear licensing from day one prevents future legal debt.

## 2. Decision Drivers

- Clarity for contributors regarding their rights.
- Protection of commercial value-adds.
- Compliance with standard OSS practices.

## 3. Consequences

Implement a dual-licensing strategy through physical repository separation.

- **OSS Repo (`dilocash-oss`):** Licensed under the permissive **MIT License**.
- **Premium Repo:** Governed by a **Proprietary Commercial License**.
- **Standard:** Every source file must include an SPDX header (e.g., `// SPDX-License-Identifier: MIT`).

- **Positive:** Transparent legal boundary for users and contributors.
- **Positive:** MIT license encourages wide adoption of the core engine.
- **Negative:** Requires legal review of contributions to ensure they are properly licensed.

---
