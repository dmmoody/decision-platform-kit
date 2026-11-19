# Spike plan

Branch: `DOC-spike-decision-kit`

## Hypothesis

A thin Go module can remove duplicated decide plumbing across domain engines **without** becoming a universal PDP service.

## In scope for the spike

- [x] Repo + module + docs stating the boundary
- [x] `decide` package: `Decide` harness with `SnapshotLoader` / `Disposer` / `Recorder`
- [x] Unit test with a toy snapshot/outcome (in-kit example, not a real product)
- [x] Semver via git tags, CHANGELOG, consumer `GOPRIVATE` docs, CI
- [x] Expand module: `decide.Engine`, `effect`, `example/toy`
- [x] Auto-version on merge to `master` (workflow + bump scripts)
- [ ] Document how a money-movement engine would plug in (adapter sketch only)
- [ ] Document how a risk-style engine would plug in (adapter sketch only)

## Out of scope for the spike

- Production positive-pay or intercept integration
- Policy authoring UI
- Expression language runtime
- Multi-tenant cloud deployment of “the kit” as a service

## Success criteria

1. A reader understands why this is not Axiomatics-for-everything
2. A Go engineer can implement a fake domain decide path using the harness in under an hour
3. No package in the kit imports a product domain
4. Module builds and tests pass with stdlib only for the spike

## Next after spike

Tag `v0.1.0` on the release commit (see [versioning.md](versioning.md)). Optionally extract helpers from `poz-pay` once the harness matches real seams (inbox/outbox, request key, effect). Aim for `v1.0.0` only when the public `decide` API is something we will not break lightly.
