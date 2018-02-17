package clock

import (
	"fmt"
	"regexp"
	"time"
)

// BrokenClock has seen better days; in its mind, time is standing still.
type BrokenClock struct {
	// T is the time that this clock always thinks it is.
	T time.Time
}

// MkBrokenClock creates a new broken clock based on a timestamp with the
// specified layout (e.g. time.RFC3339). If the timestamp is an empty string,
// the clock's time is set to the zero time (January 1, year 1, 00:00 UTC).
func MkBrokenClock(layout string, timestamp string) (BrokenClock, error) {
	if timestamp == "" {
		return BrokenClock{}, nil
	}
	var c BrokenClock
	t, err := time.Parse(layout, timestamp)
	if err != nil {
		return c, err
	}
	return BrokenClock{T: t}, nil
}

// Now always returns the same time.
func (c BrokenClock) Now() time.Time {
	return c.T
}

// MarshalJSON returns a JSON string representation of a clock.
func (c BrokenClock) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"BrokenClock{%s}"`, c.T.Format(timestampLayout))), nil
}

// UnmarshalJSON returns the clock represented by a byte array.
func (c *BrokenClock) UnmarshalJSON(b []byte) error {
	re := regexp.MustCompile(`"BrokenClock{(.*)}"`)
	timestamp := re.ReplaceAllString(string(b), "$1")
	t, err := time.Parse(timestampLayout, timestamp)
	if err != nil {
		return err
	}
	*c = BrokenClock{T: t}
	return nil
}

// Equal returns true iff a clock behaves the same as another.
func (c BrokenClock) Equal(o BrokenClock) bool {
	return c.T.Equal(o.T)
}
