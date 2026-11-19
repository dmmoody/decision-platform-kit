package toy_test

import (
	"context"
	"testing"

	"github.com/dmmoody/decision-platform-kit/decide"
	"github.com/dmmoody/decision-platform-kit/effect"
	"github.com/dmmoody/decision-platform-kit/example/toy"
)

type memList map[string]bool

func (m memList) Listed(_ context.Context, tenant, key string) (bool, error) {
	return m[tenant+"/"+key], nil
}

func TestToyEngine_allowWhenListed(t *testing.T) {
	t.Parallel()

	store := memList{"acme/item-1": true}
	eng := toy.New(effect.Authoritative, store)

	out, err := eng.Decide(context.Background(), decide.Request{
		Key:    "item-1",
		Tenant: "acme",
	})
	if err != nil {
		t.Fatalf("Decide: %v", err)
	}
	if out.Decision != "allow" {
		t.Fatalf("Decision = %q, want allow", out.Decision)
	}
	if out.Effect != effect.Authoritative {
		t.Fatalf("Effect = %q, want authoritative", out.Effect)
	}
}

func TestToyEngine_denyWhenMissing(t *testing.T) {
	t.Parallel()

	eng := toy.New(effect.Shadow, memList{})
	out, err := eng.Decide(context.Background(), decide.Request{
		Key:    "missing",
		Tenant: "acme",
	})
	if err != nil {
		t.Fatalf("Decide: %v", err)
	}
	if out.Decision != "deny" {
		t.Fatalf("Decision = %q, want deny", out.Decision)
	}
	if out.Effect != effect.Shadow {
		t.Fatalf("Effect = %q, want shadow", out.Effect)
	}
}
