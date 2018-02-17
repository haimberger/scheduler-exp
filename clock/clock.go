package clock

import "time"

const timestampLayout = time.RFC3339

// Clock keeps time.
type Clock interface {
	// Now returns the current time according to this clock.
	Now() time.Time
}

// StandardClock keeps time accurately.
type StandardClock struct{}

// Now returns the current time.
func (*StandardClock) Now() time.Time {
	return time.Now()
}

// Equal returns true iff a clock behaves the same as another.
func (*StandardClock) Equal(*StandardClock) bool {
	// Since nothing distinguishes one standard clock from another,
	// the result is always `true`.
	return true
}
