# ADR-004: Open Core Extension Model

- **Status**: Accepted
- **Date**: 2026-02-03
- **Authors**: @jalbarran
- **Technical Domain**: Backend

## 1. Context and Problem Statement

To support a sustainable business model, we want to provide an Open Source core while monetizing premium features (Advanced NLU, Bank Sync). We need to separate proprietary code without breaking the OSS build.

## 2. Decision Drivers

- Protect intellectual property of premium features.
- Maintain a functional, high-quality community version.
- Enable seamless "plug-and-play" of premium modules.

## 3. Consequences

Adopt an **Open Core** model using **Git Submodules** and Go-based **Dependency Injection**.

- Public Repo (`dilocash-oss`) contains the core engine.
- Private Repo (`dilocash-premium`) is linked as a submodule in the `/premium` directory.
- Use interface-based injection to swap implementations at runtime.

- **Positive:** Clear legal and technical boundary for IP protection.
- **Negative:** Increased complexity in CI/CD pipelines to manage private submodules.
- **Negative:** Requires strict interface discipline across the domain layer.

---
