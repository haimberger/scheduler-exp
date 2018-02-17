package clock

import "time"

// Clock keeps time.
type Clock interface {
	// Now returns the current time according to this clock.
	Now() time.Time
}
