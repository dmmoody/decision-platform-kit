// Package toy is a minimal domain decision engine that imports the kit.
//
// It is not a product. It shows how a domain supplies S and O and wires decide.Engine.
package toy

import (
	"context"

	"github.com/dmmoody/decision-platform-kit/decide"
	"github.com/dmmoody/decision-platform-kit/effect"
)

// Snapshot is the toy evaluation snapshot.
type Snapshot struct {
	Listed bool
}

// Outcome is the toy closed vocabulary.
type Outcome struct {
	Decision string
	Effect   effect.Effect
}

// Engine is a toy domain decision engine.
type Engine struct {
	inner *decide.Engine[Snapshot, Outcome]
}

// New builds a toy engine. defaultEffect is stamped on every outcome.
func New(defaultEffect effect.Effect, store ListStore) *Engine {
	d := &domain{store: store, effect: defaultEffect}
	return &Engine{
		inner: decide.New[Snapshot, Outcome](d, d, d),
	}
}

// Decide runs the kit harness for this toy domain.
func (e *Engine) Decide(ctx context.Context, req decide.Request) (Outcome, error) {
	return e.inner.Decide(ctx, req)
}

// ListStore is the toy stand-in for domain policy data (a list the domain owns).
type ListStore interface {
	Listed(ctx context.Context, tenant, key string) (bool, error)
}

type domain struct {
	store  ListStore
	effect effect.Effect
}

func (d *domain) Load(ctx context.Context, req decide.Request) (Snapshot, error) {
	ok, err := d.store.Listed(ctx, req.Tenant, req.Key)
	if err != nil {
		return Snapshot{}, err
	}
	return Snapshot{Listed: ok}, nil
}

func (d *domain) Dispose(_ context.Context, snap Snapshot) (Outcome, error) {
	decision := "deny"
	if snap.Listed {
		decision = "allow"
	}
	return Outcome{Decision: decision, Effect: d.effect}, nil
}

func (d *domain) Commit(context.Context, decide.Request, Outcome) error {
	return nil
}
