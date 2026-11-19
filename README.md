# decision-platform-kit

Go module: a **shared kit** for building domain decision engines.

This is not a hosted, company-wide decision service. It does not evaluate domain policy or enforce outcomes. It packages the chassis domain engines keep reinventing: snapshot evaluate, closed outcomes, audit/idempotency helpers, and optional shadow vs authoritative cutover.

Each domain engine owns its schemas, policy types, and outcome vocabulary. Engines written in Go import this kit.

See [docs/README.md](docs/README.md).

## Module

```text
github.com/dmmoody/decision-platform-kit
```

Tags are cut **automatically on merge to `master`**. Policy: [docs/versioning.md](docs/versioning.md). Changes: [CHANGELOG.md](CHANGELOG.md).

## Install

Private repository — set `GOPRIVATE` first:

```bash
go env -w GOPRIVATE=github.com/dmmoody/*
go get github.com/dmmoody/decision-platform-kit@latest
```

## Packages

| Import | Role |
|--------|------|
| `…/decide` | Generic load → dispose → commit harness (`Decide`, `Engine`) |
| `…/effect` | Optional `shadow` / `authoritative` stamp |
| `…/example/toy` | Minimal wiring demo (not for production use) |

```go
import (
    "github.com/dmmoody/decision-platform-kit/decide"
    "github.com/dmmoody/decision-platform-kit/effect"
)
```

## Status

Semver `v0`: API may still change. Treat `v1.0.0` as the stability promise.
