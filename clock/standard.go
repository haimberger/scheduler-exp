package clock

import "time"

// StandardClock keeps time accurately.
type StandardClock struct{}

// Now returns the current time.
func (StandardClock) Now() time.Time {
	return time.Now()
}

// MarshalJSON returns a JSON string representation of a clock.
func (StandardClock) MarshalJSON() ([]byte, error) {
	return []byte(`"StandardClock"`), nil
}

// UnmarshalJSON returns the clock represented by a byte array.
func (c *StandardClock) UnmarshalJSON([]byte) error {
	c = &StandardClock{}
	return nil
}

// Equal returns true iff a clock behaves the same as another.
func (StandardClock) Equal(StandardClock) bool {
	// Since nothing distinguishes one standard clock from another,
	// the result is always `true`.
	return true
}
