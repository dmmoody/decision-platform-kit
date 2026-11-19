package decide_test

import (
	"context"
	"testing"

	"github.com/dmmoody/decision-platform-kit/decide"
)

type snap struct {
	Match bool
}

type outcome struct {
	Decision string
}

type loader struct{}

func (loader) Load(_ context.Context, _ decide.Request) (snap, error) {
	return snap{Match: true}, nil
}

type disposer struct{}

func (disposer) Dispose(_ context.Context, s snap) (outcome, error) {
	if s.Match {
		return outcome{Decision: "allow"}, nil
	}
	return outcome{Decision: "deny"}, nil
}

type recorder struct {
	got outcome
}

func (r *recorder) Commit(_ context.Context, _ decide.Request, out outcome) error {
	r.got = out
	return nil
}

func TestDecide_toyDomain(t *testing.T) {
	t.Parallel()

	rec := &recorder{}
	out, err := decide.Decide[snap, outcome](
		context.Background(),
		loader{},
		disposer{},
		rec,
		decide.Request{Key: "k1", Tenant: "t1"},
	)
	if err != nil {
		t.Fatalf("Decide: %v", err)
	}
	if out.Decision != "allow" {
		t.Fatalf("Decision = %q, want allow", out.Decision)
	}
	if rec.got.Decision != "allow" {
		t.Fatalf("recorded = %q, want allow", rec.got.Decision)
	}
}
