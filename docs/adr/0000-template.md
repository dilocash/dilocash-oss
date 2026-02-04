To maintain the high engineering standard of **Dilocash**, every ADR should follow a consistent structure. This template is based on the **Nygard** format, which is the industry standard for capturing the "Why" behind the "What."

I recommend saving this as `docs/adr/0000-template.md`.

---

# ADR-XXXX]: [Short, Descriptive Title]

- **Status**: [Proposed | Accepted | Deprecated | Superseded by ADR-XXXX]
- **Date**: YYYY-MM-DD
- **Authors**: [@YourUsername]
- **Technical Domain**: [e.g., Backend, Frontend, Infra, AI]

## 1. Context and Problem Statement

[Describe the context of the decision. What is the specific problem we are trying to solve? Include any constraints, business requirements, or technical limitations that influenced this situation.]

## 2. Decision Drivers

- [Driver 1: e.g., We need to reduce latency for voice parsing.]
- [Driver 2: e.g., The solution must be compatible with the Open Core model.]
- [Driver 3: e.g., Minimize monthly infrastructure costs.]

## 3. Considered Options

- **Option 1**: [Brief description]
- **Option 2**: [Brief description]
- **Option 3**: [Brief description]

## 4. Decision Outcome

**Chosen Option: [Option X]**

[Explain why this option was chosen. How does it address the drivers? Why were the other options rejected?]

### Technical Implementation Details

[Optional: Provide a high-level overview of the implementation, such as a snippet of Go interfaces or a diagram of the data flow.]

## 5. Consequences

### Positive (Pros)

- [Pro 1: e.g., Improves code reusability across Web and Mobile.]
- [Pro 2: e.g., Simplifies the testing of business logic.]

### Negative (Cons/Risks)

- [Con 1: e.g., Adds complexity to the CI/CD pipeline.]
- [Con 2: e.g., Requires a specific Go version (1.22+).]

## 6. Pros and Cons of Options

### [Option 1]

- **Good**: [Benefit]
- **Bad**: [Drawback]

### [Option 2]

- **Good**: [Benefit]
- **Bad**: [Drawback]

---

### How to use this template:

1. **Numbering**: Always increment the number (`0007`, `0008`, etc.).
2. **Immutability**: Once an ADR is "Accepted" and the code is merged, **do not edit the file** to change the decision. Instead, create a new ADR that "Supersedes" the old one.
3. **Peer Review**: Treat ADRs like code. They should be submitted as Pull Requests and reviewed by teammates (or documented for your future self).
