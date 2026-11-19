// Package decide is a thin harness for domain decision engines.
//
// Domains supply snapshot type S and outcome type O. This package does not
// define business outcomes, policy stores, or enforcement actions.
package decide

import "context"

// SnapshotLoader loads an immutable evaluation snapshot for the request.
type SnapshotLoader[S any] interface {
	Load(ctx context.Context, req Request) (S, error)
}

// Disposer turns a snapshot into a domain-closed outcome.
type Disposer[S, O any] interface {
	Dispose(ctx context.Context, snap S) (O, error)
}

// Recorder persists the outcome (and domain audit or outbox) after dispose.
type Recorder[O any] interface {
	Commit(ctx context.Context, req Request, out O) error
}

// Engine binds load, dispose, and record collaborators for repeated Decide calls.
type Engine[S, O any] struct {
	load    SnapshotLoader[S]
	dispose Disposer[S, O]
	rec     Recorder[O]
}

// New constructs an Engine. All collaborators are required.
func New[S, O any](load SnapshotLoader[S], dispose Disposer[S, O], rec Recorder[O]) *Engine[S, O] {
	return &Engine[S, O]{
		load:    load,
		dispose: dispose,
		rec:     rec,
	}
}

// Decide validates the request, then load → dispose → commit.
func (e *Engine[S, O]) Decide(ctx context.Context, req Request) (O, error) {
	return run(ctx, e.load, e.dispose, e.rec, req)
}

// Decide is a one-shot harness when the caller does not keep an Engine.
// Domain gates (authorization, request-key replay, budgets) stay beside or wrap this.
func Decide[S, O any](
	ctx context.Context,
	load SnapshotLoader[S],
	dispose Disposer[S, O],
	rec Recorder[O],
	req Request,
) (O, error) {
	return run(ctx, load, dispose, rec, req)
}

func run[S, O any](
	ctx context.Context,
	load SnapshotLoader[S],
	dispose Disposer[S, O],
	rec Recorder[O],
	req Request,
) (O, error) {
	var zero O

	if err := req.Validate(); err != nil {
		return zero, err
	}

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
