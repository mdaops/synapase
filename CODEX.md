## Codex

This products goal is to ease development of building out a platform for engineering teams.
The goal, should be to reduce the time it takes to make a developer platform. Rather than going from 0 you started from 70.

The project contains these core components:

- `packages`: Control plane elements installed on the consumer cluster.
- `apis`: The product APIs for creating and managing products
- `controllers`: custom Kubernetes operators (e.g. k6-operator, entitlements)
- `config`: Kustomize-managed YAML for Crossplane, Flux, and Kyverno
- `cluster`: Cluster-level resources (e.g. RBAC, NetworkPolicy, CRDs for managed resources)
- `examples`: sample overlays for AWS and GCP
- `scripts`: helper scripts for codegen and manifest generation
- `docs`: documentation and architecture notes
- `.github/workflows`: CI pipelines

