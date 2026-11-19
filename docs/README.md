# Decision platform kit — docs

## Why this exists

Organizations want reuse so every business domain does not reinvent a decision engine. A single company-wide decision *service* for every question becomes a junk drawer: wrong outcome vocabularies, attribute fetches into every database, and overhead when the caller already knows the answer.

This module takes the middle path:

| Avoid | Prefer |
|-------|--------|
| One universal decide service for all domains | Domain decision **engines** (services or libraries) |
| Copy-paste evaluate / audit / cutover plumbing | Shared **kit** (this module) |
| One shared outcome enum for the company | Domain-closed outcomes |
| Kit owns every domain list or attribute store | Domain owns its data; kit owns evaluate plumbing |

Reuse the **chassis**. Do not force one engine for every domain question.

## What this kit is

A Go library that standardizes:

1. **Decide harness** — load snapshot → dispose → commit/record
2. **Request identity** — key + tenant (idempotency hooks belong beside or wrap the harness)
3. **Optional cutover stamp** — shadow vs authoritative (`effect`)
4. **Audit-shaped recording** — outcome persistence the domain defines
5. **Docs and examples** — how a domain plugs in

It is the paved road for building domain decision capabilities in Go—not a general policy-language platform.

## What this kit is not

- Not a hosted decision service all apps call
- Not a general policy language product (Rego/XACML/Cedar as the platform)
- Not a store for domain lists, signals, or attributes
- Not a replacement for any domain product’s decision service

Domain engines may import this kit (when they are Go) or follow the same pattern in another language. They do not collapse into one binary.

## PDP/PEP — pattern only

- **PEP (pattern):** the caller that asks for a decision and then enforces.
- **PDP (pattern):** the domain engine that answers for **its** question.

This kit helps you build domain PDPs consistently. It is not “the” corporate PDP.

## Further reading

- Spike plan: [spike.md](spike.md)
- Architecture: [architecture.md](architecture.md)
- Versioning: [versioning.md](versioning.md)
