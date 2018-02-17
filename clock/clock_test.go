package clock

import (
	"encoding/json"
	"testing"
	"time"
)

func TestStandardNow(t *testing.T) {
	var c Clock = &StandardClock{} // make sure that StandardClock implements Clock
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
	expected := "{}"
	c := StandardClock{}
	actual, err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
	} else if string(actual) != expected {
		t.Fatalf("expected %s; got %s", expected, actual)
	}
}

func TestStandardUnmarshalJSON(t *testing.T) {
	type testCase struct {
		name    string
		jsonStr string
		err     bool
	}
	tcs := []testCase{
		{name: "typical", jsonStr: `{}`},
		{name: "ignore", jsonStr: `{"time":"2018-02-16T18:21:34Z"}`},
		{name: "invalid", jsonStr: `""`, err: true},
	}
	for _, tc := range tcs {
		var c StandardClock
		expected := StandardClock{}
		err := json.Unmarshal([]byte(tc.jsonStr), &c)
		if err == nil && tc.err {
			t.Fatalf("[%s] expected an error, but didn't get one; got %v instead", tc.name, c)
		} else if err != nil && !tc.err {
			t.Fatalf("[%s] got error: %v", tc.name, err)
		} else if err != nil {
			t.Logf("[%s] got expected error: %v", tc.name, err)
		} else if err == nil && !c.Equal(&expected) {
			t.Fatalf("[%s] expected %v; got %v", tc.name, expected, c)
		}
	}
}
