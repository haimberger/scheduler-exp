package clock

import "time"

// BrokenClock has seen better days; in its mind, time is standing still.
type BrokenClock struct {
	// T is the time that this clock always thinks it is.
	T time.Time `json:"time"`
}

// TestClock creates a new broken clock based on a timestamp with the
// specified layout (e.g. time.RFC3339). If the timestamp is an empty string,
// the clock's time is set to the zero time (January 1, year 1, 00:00 UTC).
func TestClock(layout string, timestamp string) (*BrokenClock, error) {
	if timestamp == "" {
		return &BrokenClock{}, nil
	}
	t, err := time.Parse(layout, timestamp)
	if err != nil {
		return nil, err
	}
	return &BrokenClock{T: t}, nil
}

// Now always returns the same time.
func (c *BrokenClock) Now() time.Time {
	return c.T
}

// Equal returns true iff a clock behaves the same as another.
func (c *BrokenClock) Equal(o *BrokenClock) bool {
	return c.T.Equal(o.T)
}
