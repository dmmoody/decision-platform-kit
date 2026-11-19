# Decision platform kit — docs

## Why this exists

The org wants reuse so every business domain does not reinvent a “decision engine.” A single company-wide PDP/PEP service for every question (payments, fraud, IP lists, IAM) becomes a junk drawer: wrong outcome vocabularies, PIPs into every database, and overhead when the caller already knows the answer.

This module takes the middle path:

| Avoid | Prefer |
|-------|--------|
| One universal decide service for all domains | Domain decision **services** (or modules) |
| Copy-paste inbox/audit/shadow every time | Shared **kit** (this repo) |
| Shared outcome enum for the company | Domain-closed outcomes (`pay\|return\|review`, `Allow\|Deny`, …) |
| PDP owns every list (registers, IPs, …) | Domain owns its data; kit owns evaluate plumbing |

Reuse the **chassis**. Do not force one car for every road.

## What this kit is

A Go library (and, as the spike matures, small helpers) that standardizes:

1. **Decide harness** — load snapshot → dispose → commit/record
2. **Request identity / idempotency hooks** — same key, safe retry
3. **Shadow vs enforce** — record without acting as a first-class effect
4. **Audit-shaped recording** — outcome + policy versions consulted (domain fills fields)
5. **Docs and examples** — how a domain plugs in

It is the paved road for “how JH (or we) build a domain decision capability,” not Axiomatics-in-a-box.

## What this kit is not

- Not a hosted PDP all apps call
- Not a general policy language product (Rego/XACML/Cedar as the platform)
- Not a store for check registers, risk signals, or IP allowlists
- Not a replacement for jh-intercept or a money-movement decision service

Those systems may **use** patterns from this kit (or later import packages). They do not collapse into one binary.

## PDP/PEP — pattern only

Use the words carefully:

- **PEP (pattern):** the caller that asks for a decision and then enforces (post, deny, hold, allow traffic).
- **PDP (pattern):** the domain engine that answers for **its** question.

This kit helps you build domain PDPs consistently. It is not “the” PDP.

## Spike goals

See [spike.md](spike.md). Architecture intent: [architecture.md](architecture.md). Versioning: [versioning.md](versioning.md).

## Related product docs (out of tree)

- Money-movement decision brief: `poz-pay` → `docs/Decision-Engine-Design-Brief.md`
- Risk evaluation platform: `jh-intercept` (own control/data plane; sibling, not a consumer mandate yet)
