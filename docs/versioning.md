# Versioning

This module follows [Go modules](https://go.dev/ref/mod) and [semantic versioning](https://semver.org/) via **git tags**. There is no `VERSION` file; `go get` resolves tags from the repository.

## Compatibility policy

| Range | Promise |
|-------|---------|
| **`v0.x.x`** | Spike / early API. Breaking changes allowed in minor or patch as needed. Prefer minor bumps for visible API breaks. |
| **`v1.x.x`** | Stable public API for packages we document as public. Breaking changes require `v2` (and `module …/v2` in `go.mod`). |
| **Pre-releases** | Optional: `v0.2.0-alpha.1`. Consumers must ask for them explicitly. |

Until the first `v1.0.0`, treat every import as a deliberate spike dependency.

## Tagging a release

From a clean `master` (or the commit you intend to publish):

```bash
# CHANGELOG updated for the version first, then:
git tag -a v0.1.0 -m "v0.1.0"
git push origin v0.1.0
```

Optional: create a GitHub Release from that tag so humans see notes next to the tag.

Do **not** move or retag a published `vX.Y.Z` that others may have fetched. Cut `vX.Y.Z+1` instead.

## Consuming (private module)

This repository is private. Consumers must bypass the public module proxy for this path:

```bash
go env -w GOPRIVATE=github.com/dmmoody/*
```

Then:

```bash
go get github.com/dmmoody/decision-platform-kit@v0.1.0
```

Or in `go.mod`:

```go
require github.com/dmmoody/decision-platform-kit v0.1.0
```

SSH or a GitHub credential that can read the repo is required (`git` must clone `github.com/dmmoody/decision-platform-kit`).

## Public API surface (v0)

Intended import paths:

- `github.com/dmmoody/decision-platform-kit/decide`

Anything under `internal/` (when added) is not importable by other modules. Docs-only trees are not Go packages.

## Checklist before cutting a tag

1. `go test ./...` passes
2. [CHANGELOG.md](../CHANGELOG.md) has an entry for the version
3. Tag message matches the version (`v0.1.0`)
4. Push the tag (humans); CI should already be green on the commit
