package clock

import "time"

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

// BrokenClock has seen better days; in its mind, time is standing still.
type BrokenClock struct {
	// T is the time that this clock always thinks it is.
	T time.Time
}

// Now always returns the same time.
func (c *BrokenClock) Now() time.Time {
	return c.T
}
