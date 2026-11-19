// Package effect defines the shadow vs authoritative stamp for a recorded decision.
//
// The kit records effect; callers that enforce choose whether to act on it.
// Domains that do not need cutover may ignore this package.
package effect

import "fmt"

// Effect is whether irreversible side effects may follow the outcome.
type Effect string

const (
	// Shadow means the decision is real and recorded, but enforcers must not act on it.
	Shadow Effect = "shadow"
	// Authoritative means enforcers may act on the outcome.
	Authoritative Effect = "authoritative"
)

// Valid reports whether e is a known effect value.
func (e Effect) Valid() bool {
	switch e {
	case Shadow, Authoritative:
		return true
	default:
		return false
	}
}

// Parse accepts a wire string and returns a typed Effect.
func Parse(s string) (Effect, error) {
	e := Effect(s)
	if !e.Valid() {
		return "", fmt.Errorf("effect: unknown value %q", s)
	}
	return e, nil
}
