package clock

import "time"

const timestampLayout = time.RFC3339

// Clock keeps time.
type Clock interface {
	// Now returns the current time according to this clock.
	Now() time.Time
}
