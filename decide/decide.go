// Package decide is a thin harness for domain decision engines.
//
// Domains supply snapshot type S and outcome type O. This package does not
// know payment rails, fraud signals, or IAM resources.
package decide

import "context"

// Request is the kit-level identity for one decide attempt.
// Domains may embed richer envelopes; the harness only needs these fields.
type Request struct {
	// Key is the caller-chosen idempotency key for this attempt.
	Key string
	// Tenant is an opaque tenant stamp (institution id, tenant path, …).
	Tenant string
}

// SnapshotLoader loads an immutable evaluation snapshot for the request.
type SnapshotLoader[S any] interface {
	Load(ctx context.Context, req Request) (S, error)
}

// Disposer turns a snapshot into a domain-closed outcome.
type Disposer[S, O any] interface {
	Dispose(ctx context.Context, snap S) (O, error)
}

// Recorder persists the outcome (and domain audit/outbox) after dispose.
type Recorder[O any] interface {
	Commit(ctx context.Context, req Request, out O) error
}

// Decide loads a snapshot, disposes once, then commits the outcome.
// Domain gates (permits, request-key replay, budgets) stay beside or wrap this.
func Decide[S, O any](
	ctx context.Context,
	load SnapshotLoader[S],
	dispose Disposer[S, O],
	rec Recorder[O],
	req Request,
) (O, error) {
	var zero O

	snap, err := load.Load(ctx, req)
	if err != nil {
		return zero, err
	}

	out, err := dispose.Dispose(ctx, snap)
	if err != nil {
		return zero, err
	}

	if err := rec.Commit(ctx, req, out); err != nil {
		return zero, err
	}

	return out, nil
}
