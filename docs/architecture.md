# Architecture

## Shape

```text
┌──────────────────────────────────────────┐
│  decision-platform-kit (Go module)       │
│  harness · identity · shadow · record    │
└──────────────────────────────────────────┘
                 ▲
                 │ import
     ┌───────────┴───────────┐
     │                       │
┌────┴─────────┐    ┌────────┴─────────┐
│ Domain A     │    │ Domain B         │
│ e.g. money   │    │ e.g. risk        │
│ movement     │    │                  │
│ outcomes +   │    │ outcomes +       │
│ policy types │    │ policy types     │
└──────────────┘    └──────────────────┘
```

Each domain service remains the system of record for its policy data and outcome vocabulary. The kit supplies shared evaluate plumbing.

## Go approach (no class hierarchy)

Variation is interfaces and generics, not `AbstractDecisionEngine` subclasses.

Public packages (v0; names may still evolve):

| Package | Role |
|---------|------|
| `decide` | `Request`, `SnapshotLoader`, `Disposer`, `Recorder`, `Decide`, `Engine` |
| `effect` | `shadow` / `authoritative` |
| `example/toy` | In-repo wiring demo only |

Domain packages own snapshot `S` and outcome `O`. The kit never imports product domains.

## Effect / shadow

Where a domain needs cutover without dual brains, the kit will treat **effect** (e.g. `shadow` vs `authoritative`) as a recorded stamp the **caller** respects. The kit does not call cores or fraud actions.

## Deliberate non-goals

- Cross-domain decide API
- Shared PIP registry into arbitrary databases
- Company-wide allow/deny as the only outcome
- Forcing one expression language on every domain in v0

## Evolution

1. Spike harness + docs (this branch)
2. Idempotency and rejection shapes that domains can adopt
3. Optional outbox helper
4. Example domain stub in-repo (toy, not production positive pay)
5. Extract only what a second real consumer needs

Copy until the second consumer; then harden the kit API. Versions are git tags (`v0.x` while spiking); see [versioning.md](versioning.md).
