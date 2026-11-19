# decision-platform-kit

Go module: a **shared kit** for building domain decision engines.

This is not a universal PDP service. It does not decide payments, fraud, or IAM in one place. It packages the chassis domains keep reinventing: snapshot evaluate, closed outcomes, audit/idempotency helpers, and shadow vs enforce plumbing.

Domain products (positive pay, risk/intercept-style engines, and others) own their schemas, policy types, and outcome vocabulary. They import this kit.

See [docs/README.md](docs/README.md).

## Module

```text
github.com/dmmoody/decision-platform-kit
```

## Status

Spike. APIs will change. Do not treat the surface as stable until a tagged release says so.
