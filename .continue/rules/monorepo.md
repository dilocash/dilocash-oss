---
name: Monorepo Architecture Rules
globs: ["**/*"]
---

# Dilocash Monorepo Guidelines

You are an expert architect for the Dilocash project. This is a monorepo containing multiple Go packages and infrastructure code. Follow these rules when generating or modifying code:

## 1. Package Boundaries
- **`dilocash-oss/`**: Core ledger logic, gRPC API definitions (v1), and open-source entities. No premium features allowed here.
- **`dilocash-premium/`**: Proprietary extensions. Can import `dilocash-oss` but `dilocash-oss` must NEVER import from here.
- **`dilocash-infra/`**: Pulumi code for AWS (Fargate/S3). Should not contain business logic.

# Dependency Management
- Use **Go Modules** within each package.
- When adding dependencies, ensure they are compatible with the **MIT License** for `dilocash-oss`.
- Prefer **Connect-Go** over standard `grpc-go` for API implementations to support Web/Mobile/Hooks on a single port (ADR-041).

# Communication
- All internal service communication must use **gRPC/Protobuf**.
- Always check the `proto/` directory and all subdirectories under it. for existing service contracts before creating new ones.

# Monorepo Map for Agent
- gRPC definitions are in `proto/dilocash/v1/` directory and all subdirectories under it.
- Go handlers are in `apps/api/internal/generated/transport`.
- Generated sqlc code is in `apps/api/internal/generated/db`.
- Use the `read_file` and `search_code` tools to explore these directories before asking the user for code.