package effect_test

import (
	"testing"

	"github.com/dmmoody/decision-platform-kit/effect"
)

func TestParse(t *testing.T) {
	t.Parallel()

	e, err := effect.Parse("shadow")
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if e != effect.Shadow {
		t.Fatalf("got %q, want shadow", e)
	}

	if _, err := effect.Parse("nope"); err == nil {
		t.Fatal("expected error")
	}
}

func TestValid(t *testing.T) {
	t.Parallel()

	if !effect.Authoritative.Valid() {
		t.Fatal("authoritative should be valid")
	}
	if effect.Effect("x").Valid() {
		t.Fatal("x should be invalid")
	}
}
