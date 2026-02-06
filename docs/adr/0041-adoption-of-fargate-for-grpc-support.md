# ADR-041: Adoption of AWS Fargate for gRPC Support

- **Status**: Accepted
- **Date**: 2026-02-06
- **Authors**: @jalbarran
- **Technical Domain**: Infrastructure / Backend

## 1. Context and Problem Statement

The Dilocash backend uses **gRPC** as its primary communication protocol (ADR-030). We initially considered **AWS App Runner** for its simplicity and "no-ops" experience. However, AWS App Runner has significant limitations regarding gRPC support:

1. It does not fully support HTTP/2 trailers, which are essential for gRPC status codes.
2. It lacks support for gRPC-standard load balancing and health checks in its current managed state.
3. Bidirectional streaming and specific gRPC features are unreliable on the platform.

To maintain a performant and reliable gRPC API, we need a platform that provides full HTTP/2 support and integration with an Application Load Balancer (ALB) configured for gRPC.

## 2. Decision Drivers

- **Requirement for full gRPC support:** The core communication depends on gRPC (ADR-012, ADR-030).
- **Control over Networking:** Ability to configure ALB for gRPC (Protocol version HTTP2, gRPC health checks).
- **Scalability:** The solution must scale horizontally to handle varying traffic.
- **Cost-Benefit:** While Fargate is slightly more complex to manage than App Runner, it provides the necessary technical capabilities without the overhead of managing EC2 instances.

## 3. Consequences

### The Decision

We will use **AWS ECS on Fargate** as the deployment target for the Dilocash Go API.

- **Load Balancing:** Use an **Application Load Balancer (ALB)** with a listener configured for HTTPS/HTTP2 and a target group using the `GRPC` protocol.
- **Infrastructure as Code:** Implement the Fargate service, Task Definitions, and ALB using **Pulumi (Go)** (ADR-023).
- **Service Mesh:** Consider AWS App Mesh or standard ECS Service Connect for internal service-to-service communication in the future.

### Implications

- **Positive:** Full compatibility with gRPC and HTTP/2.
- **Positive:** Granular control over CPU, memory, and networking.
- **Positive:** Better integration with AWS CloudWatch for structured logging and metrics.
- **Negative:** Increased complexity in infrastructure setup compared to App Runner (requires VPC, Subnets, Security Groups, IAM Roles, etc.).
- **Negative:** Higher initial management overhead for the Pulumi IaC.
