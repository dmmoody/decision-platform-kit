# Architecture

## Shape

```text
┌──────────────────────────────────────────┐
│  decision-platform-kit (Go module)       │
│  harness · identity · effect · record    │
└──────────────────────────────────────────┘
                 ▲
                 │ import
     ┌───────────┴───────────┐
     │                       │
┌────┴─────────┐    ┌────────┴─────────┐
│ Domain A     │    │ Domain B         │
│ engine       │    │ engine           │
│ outcomes +   │    │ outcomes +       │
│ policy types │    │ policy types     │
└──────────────┘    └──────────────────┘
```

Each domain engine remains the system of record for its policy data and outcome vocabulary. The kit supplies shared evaluate plumbing.

## Go approach (no class hierarchy)

Variation is interfaces and generics, not inheritance.

Public packages (v0; names may still evolve):

| Package | Role |
|---------|------|
| `decide` | `Request`, `SnapshotLoader`, `Disposer`, `Recorder`, `Decide`, `Engine` |
| `effect` | Optional `shadow` / `authoritative` |
| `example/toy` | In-repo wiring demo only |

Domain packages own snapshot `S` and outcome `O`. The kit never imports product domains.

## Effect / shadow

Where a domain needs cutover without dual brains, **effect** (`shadow` vs `authoritative`) is a recorded stamp the **caller** respects. The kit does not perform enforcement actions.

## Deliberate non-goals

- Cross-domain decide API
- Shared attribute/PIP registry into arbitrary databases
- Company-wide allow/deny as the only outcome
- Forcing one expression language on every domain in v0

## Evolution

1. Harness + docs
2. Idempotency and rejection shapes domains can adopt
3. Optional outbox helper
4. Example domain stub in-repo (toy only)
5. Extract only what a second real consumer needs

Copy until the second consumer; then harden the kit API. Versions are git tags (`v0.x` while early); see [versioning.md](versioning.md).
