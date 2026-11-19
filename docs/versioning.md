# Versioning

This module follows [Go modules](https://go.dev/ref/mod) and [semantic versioning](https://semver.org/) via **git tags**. There is no `VERSION` file; `go get` resolves tags from the repository.

## Auto-release on merge

Every merge (push) to `master` runs [`.github/workflows/release.yml`](../.github/workflows/release.yml):

1. `go test ./...`
2. Compute the next `vX.Y.Z` from commits since the latest `v*` tag
3. Create an annotated git tag
4. Create a GitHub Release with the commit list

Put **`[skip release]`** in the merge commit subject to skip tagging (rare).

### Bump rules (`TYPE:` prefixes)

| Commit prefix | Bump |
|---------------|------|
| `FEAT:`, `ARCH:`, `REFACT:` | minor |
| `FIX:`, `PERF:`, `DOC:`, `CHORE:`, `TEST:`, `STYLE:`, `CI:` | patch |
| `BREAKING:`, `MAJOR:` | major — but while still on **`v0`**, treated as **minor** (no silent jump to `v1`) |
| no recognized prefix | patch |

Examples after `v0.1.0`:

- merge with `FEAT: …` → `v0.2.0`
- merge with only `FIX: …` / `DOC: …` → `v0.1.1`

Keep [CHANGELOG.md](../CHANGELOG.md) updated in the PR under **Unreleased**; after the auto-tag, fold that section into the new version in a follow-up if you want the file to match tags exactly. The GitHub Release notes are the machine-generated source of truth for what landed in the tag.

## Compatibility policy

| Range | Promise |
|-------|---------|
| **`v0.x.x`** | Spike / early API. Breaking changes allowed. Prefer `FEAT`/`ARCH` for visible API breaks. |
| **`v1.x.x`** | Stable public API. Breaking changes require `v2` (and `module …/v2` in `go.mod`). |
| **Pre-releases** | Not produced by the auto-release workflow. Cut manually if needed (`v0.2.0-alpha.1`). |

Until the first `v1.0.0`, treat every import as a deliberate spike dependency.

## Manual tag (escape hatch)

Only if automation cannot run:

```bash
./scripts/next-version.sh    # preview
git tag -a v0.2.0 -m "v0.2.0"
git push origin v0.2.0
gh release create v0.2.0 --generate-notes
```

Do **not** move or retag a published `vX.Y.Z`. Cut the next version instead.

## Consuming (private module)

```bash
go env -w GOPRIVATE=github.com/dmmoody/*
go get github.com/dmmoody/decision-platform-kit@v0.2.0   # or @latest
```

```go
require github.com/dmmoody/decision-platform-kit v0.2.0
```

SSH or a GitHub credential that can read the repo is required.

## Public API surface (v0)

Intended import paths:

- `github.com/dmmoody/decision-platform-kit/decide`
- `github.com/dmmoody/decision-platform-kit/effect`

`example/toy` is demonstrative only. Anything under `internal/` (when added) is not importable by other modules.
