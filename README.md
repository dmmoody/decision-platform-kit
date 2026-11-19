# decision-platform-kit

Go module: a **shared kit** for building domain decision engines.

This is not a universal PDP service. It does not decide payments, fraud, or IAM in one place. It packages the chassis domains keep reinventing: snapshot evaluate, closed outcomes, audit/idempotency helpers, and shadow vs enforce plumbing.

Domain products (positive pay, risk/intercept-style engines, and others) own their schemas, policy types, and outcome vocabulary. They import this kit.

See [docs/README.md](docs/README.md).

## Module

```text
github.com/dmmoody/decision-platform-kit
```

Current release tag: **`v0.1.0`** (`v0` = spike; breaking changes allowed). Versioning: [docs/versioning.md](docs/versioning.md). Changes: [CHANGELOG.md](CHANGELOG.md).

## Install

Private repo — set `GOPRIVATE` first:

```bash
go env -w GOPRIVATE=github.com/dmmoody/*
go get github.com/dmmoody/decision-platform-kit@v0.1.0
```

## Packages

| Import | Role |
|--------|------|
| `…/decide` | Generic load → dispose → commit harness (`Decide`, `Engine`) |
| `…/effect` | `shadow` / `authoritative` stamp |
| `…/example/toy` | Minimal domain showing how to wire the kit (not for production) |

```go
import (
    "github.com/dmmoody/decision-platform-kit/decide"
    "github.com/dmmoody/decision-platform-kit/effect"
)
```

## Status

Spike under semver `v0`. Do not assume API stability until `v1.0.0`.
