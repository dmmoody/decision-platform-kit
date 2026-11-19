# Spike plan

## Hypothesis

A thin Go module can remove duplicated decide plumbing across domain engines **without** becoming a universal hosted decision service.

## In scope for the spike

- [x] Repo + module + docs stating the boundary
- [x] `decide` package: `Decide` harness with `SnapshotLoader` / `Disposer` / `Recorder`
- [x] Unit test with a toy snapshot/outcome (in-kit example, not a real product)
- [x] Semver via git tags, CHANGELOG, consumer `GOPRIVATE` docs, CI
- [x] Expand module: `decide.Engine`, `effect`, `example/toy`
- [x] Auto-version on merge to `master` (workflow + bump scripts)
- [x] Domain-neutral public docs and package comments
- [ ] Adapter sketches for two example domains (optional teaching docs)

## Out of scope for the spike

- Production integration into any specific product
- Policy authoring UI
- Expression language runtime
- Multi-tenant cloud deployment of “the kit” as a service

## Success criteria

1. A reader understands why this is not a company-wide decision service
2. A Go engineer can implement a fake domain decide path using the harness in under an hour
3. No package in the kit imports a product domain
4. Module builds and tests pass with stdlib only for the spike
5. Public godoc and README do not bind the kit to one business domain

## Next after spike

Live on `v0.x` with real Go consumers. Extract helpers (request-key replay, inbox/outbox) only when a second consumer needs them. Aim for `v1.0.0` when the public `decide` API is something we will not break lightly.
