package decide_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dmmoody/decision-platform-kit/decide"
)

type snap struct {
	Match bool
}

type outcome struct {
	Decision string
}

type loader struct {
	s   snap
	err error
}

func (l loader) Load(_ context.Context, _ decide.Request) (snap, error) {
	return l.s, l.err
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
	err error
}

func (r *recorder) Commit(_ context.Context, _ decide.Request, out outcome) error {
	if r.err != nil {
		return r.err
	}
	r.got = out
	return nil
}

func TestDecide_toyDomain(t *testing.T) {
	t.Parallel()

	rec := &recorder{}
	out, err := decide.Decide[snap, outcome](
		context.Background(),
		loader{s: snap{Match: true}},
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

func TestEngine_Decide(t *testing.T) {
	t.Parallel()

	rec := &recorder{}
	eng := decide.New[snap, outcome](loader{s: snap{Match: false}}, disposer{}, rec)
	out, err := eng.Decide(context.Background(), decide.Request{Key: "k2", Tenant: "t1"})
	if err != nil {
		t.Fatalf("Decide: %v", err)
	}
	if out.Decision != "deny" {
		t.Fatalf("Decision = %q, want deny", out.Decision)
	}
}

func TestDecide_rejectsEmptyRequest(t *testing.T) {
	t.Parallel()

	_, err := decide.Decide[snap, outcome](
		context.Background(),
		loader{},
		disposer{},
		&recorder{},
		decide.Request{},
	)
	if err == nil {
		t.Fatal("expected validation error")
	}
}

func TestDecide_loadError(t *testing.T) {
	t.Parallel()

	want := errors.New("load failed")
	_, err := decide.Decide[snap, outcome](
		context.Background(),
		loader{err: want},
		disposer{},
		&recorder{},
		decide.Request{Key: "k", Tenant: "t"},
	)
	if !errors.Is(err, want) {
		t.Fatalf("err = %v, want %v", err, want)
	}
}
