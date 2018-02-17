package clock

import (
	"testing"
	"time"
)

func TestBrokenClock(t *testing.T) {
	var zeroTime time.Time
	testCases := []time.Time{zeroTime, time.Now()}
	for _, tc := range testCases {
		c := BrokenClock{T: tc}
		actual := c.Now()
		if !actual.Equal(tc) {
			t.Fatalf("expected %v; got %v", tc, actual)
		}
	}
}
