package clock

import "time"

// StandardClock keeps time accurately.
type StandardClock struct{}

// Now returns the current time.
func (StandardClock) Now() time.Time {
	return time.Now()
}
