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

Typical spike surface (names will evolve):

| Symbol | Role |
|--------|------|
| `Request` | Caller identity for this decide (key, tenant scope stamps) |
| `SnapshotLoader[S]` | Domain loads immutable snapshot `S` for the request |
| `Disposer[S, O]` | Domain turns snapshot into closed outcome `O` |
| `Recorder[O]` | Domain commits outcome + audit/outbox as it defines |
| `Decide` | Kit orchestrates load → dispose → record (and later gates) |

Domain packages own `S` and `O`. The kit never imports domain types.

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

Copy until the second consumer; then harden the kit API and tag a module version.
