package clock

import "time"

// BrokenClock has seen better days; in its mind, time is standing still.
type BrokenClock struct {
	// T is the time that this clock always thinks it is.
	T time.Time
}

// Now always returns the same time.
func (c *BrokenClock) Now() time.Time {
	return c.T
}
