# Changelog

All notable changes to this module are documented here.
Versions match git tags (`vMAJOR.MINOR.PATCH`).

## [Unreleased]

### Added

- `decide.Engine` with constructor `New` and request validation
- `effect` package (`shadow` / `authoritative`)
- `example/toy` domain showing kit wiring
- Auto-release on merge to `master` (semver tag + GitHub Release)

## [v0.1.0] - 2026-11-19

### Added

- Module `github.com/dmmoody/decision-platform-kit`
- `decide` package: generic `Decide` harness with `SnapshotLoader`, `Disposer`, `Recorder`
- Docs: kit boundary, architecture, spike plan, versioning
- CI workflow: `go test ./...` on push and pull request

[Unreleased]: https://github.com/dmmoody/decision-platform-kit/compare/v0.1.0...HEAD
[v0.1.0]: https://github.com/dmmoody/decision-platform-kit/releases/tag/v0.1.0
