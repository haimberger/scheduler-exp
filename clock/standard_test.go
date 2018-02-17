package clock

import (
	"encoding/json"
	"testing"
	"time"
)

func TestStandardNow(t *testing.T) {
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

func TestStandardMarshalJSON(t *testing.T) {
	expected := `"StandardClock"`
	c := StandardClock{}
	actual, err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
	} else if string(actual) != expected {
		t.Fatalf("expected %s; got %s", expected, actual)
	}
}

func TestStandardUnmarshalJSON(t *testing.T) {
	var c StandardClock
	if err := json.Unmarshal([]byte(`"StandardClock"`), &c); err != nil {
		t.Fatal(err)
	}
	expected := StandardClock{}
	if !c.Equal(expected) {
		t.Fatalf("expected %v; got %v", c, expected)
	}
}
