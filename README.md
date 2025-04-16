 # Control Plane Monorepo

 This repository contains:

- `packages`: Go-based Crossplane packages (providers, compositions, claims)
 - `controllers`: custom Kubernetes operators (e.g. k6-operator)
 - `config`: Kustomize-managed YAML for Crossplane, Flux, and Kyverno
 - `examples`: sample overlays for AWS and GCP
 - `scripts`: helper scripts for codegen and manifest generation
 - `docs`: documentation and architecture notes
 - `.github/workflows`: CI pipelines

 Initialize modules with:

 ```bash
 go mod init github.com/<your-org>/control-plane
 ```
