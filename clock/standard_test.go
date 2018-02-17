package clock

import (
	"testing"
	"time"
)

func TestStandardClock(t *testing.T) {
	c := StandardClock{}
	before := time.Now()
	actual := c.Now()
	after := time.Now()
	if !before.Before(after) {
		t.Fatal("time is going backwards?!")
	} else if !before.Before(actual) {
		t.Fatalf("expected %v to be before %v; it wasn't", before, actual)
	} else if !after.After(actual) {
		t.Fatalf("expected %v to be after %v; it wasn't", after, actual)
	}
}
