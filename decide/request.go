package decide

import "fmt"

// Request is the kit-level identity for one decide attempt.
// Domains may carry richer envelopes beside this; the harness only requires these fields.
type Request struct {
	// Key is the caller-chosen idempotency key for this attempt.
	Key string
	// Tenant is an opaque tenant stamp (institution id, tenant path, …).
	Tenant string
}

// Validate reports whether Key and Tenant are present.
func (r Request) Validate() error {
	if r.Key == "" {
		return fmt.Errorf("decide: request key is required")
	}
	if r.Tenant == "" {
		return fmt.Errorf("decide: request tenant is required")
	}
	return nil
}
