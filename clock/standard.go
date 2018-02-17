package clock

import "time"

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
