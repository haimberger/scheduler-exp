package clock

import (
	"encoding/json"
	"testing"
)

func TestMkBrokenClock(t *testing.T) {
	type testCase struct {
		name      string
		layout    string
		timestamp string
		err       bool
	}
	tcs := []testCase{
		{name: "no timestamp", layout: "", timestamp: "", err: false},
		{name: "no layout", layout: "", timestamp: "2018-02-16T18:21:34Z", err: true},
		{name: "mismatched layout", layout: "Mon Jan 2 15:04:05 MST 2006", timestamp: "2018-02-16T18:21:34Z", err: true},
		{name: "matching layout", layout: timestampLayout, timestamp: "2018-02-16T18:21:34Z", err: false},
	}
	for _, tc := range tcs {
		_, err := TestClock(tc.layout, tc.timestamp)
		if err == nil && tc.err {
			t.Fatalf("[%s] expected an error, but didn't get one", tc.name)
		} else if err != nil && !tc.err {
			t.Fatalf("[%s] got error: %v", tc.name, err)
		}
	}
}

func TestBrokenNow(t *testing.T) {
	type testCase struct {
		name      string
		timestamp string
		expected  string
	}
	tcs := []testCase{
		{name: "default", timestamp: "", expected: "0001-01-01T00:00:00Z"},
		{name: "typical", timestamp: "2018-02-16T18:21:34Z", expected: "2018-02-16T18:21:34Z"},
	}
	for _, tc := range tcs {
		var c Clock // make sure that BrokenClock implements Clock
		c, err := TestClock(timestampLayout, tc.timestamp)
		if err != nil {
			t.Fatalf("[%s] got error: %v", tc.name, err)
		}
		actual := c.Now().Format(timestampLayout)
		if actual != tc.expected {
			t.Fatalf("[%s] expected %v; got %v", tc.name, tc.expected, actual)
		}
	}
}

func TestBrokenMarshalJSON(t *testing.T) {
	type testCase struct {
		name      string
		timestamp string
		expected  string
	}
	tcs := []testCase{
		{name: "default", timestamp: "", expected: `{"time":"0001-01-01T00:00:00Z"}`},
		{name: "typical", timestamp: "2018-02-16T18:21:34Z", expected: `{"time":"2018-02-16T18:21:34Z"}`},
	}
	for _, tc := range tcs {
		c, err := TestClock(timestampLayout, tc.timestamp)
		if err != nil {
			t.Fatalf("[%s] got error: %v", tc.name, err)
		}
		actual, err := json.Marshal(c)
		if err != nil {
			t.Fatalf("[%s] got error: %v", tc.name, err)
		} else if string(actual) != tc.expected {
			t.Fatalf("[%s] expected %s; got %s", tc.name, tc.expected, actual)
		}
	}
}

func TestBrokenUnmarshalJSON(t *testing.T) {
	type testCase struct {
		name      string
		jsonStr   string
		timestamp string
		err       bool
	}
	tcs := []testCase{
		{name: "empty", jsonStr: `{}`, timestamp: ""},
		{name: "ignore", jsonStr: `{"foo":"bar"}`, timestamp: ""},
		{name: "default", jsonStr: `{"time":"0001-01-01T00:00:00Z"}`, timestamp: ""},
		{name: "typical", jsonStr: `{"time":"2018-02-16T18:21:34Z"}`, timestamp: "2018-02-16T18:21:34Z"},
		{name: "invalid", jsonStr: `{"time":"02/16/2018 18:21:34"}`, timestamp: "", err: true},
	}
	for _, tc := range tcs {
		expected, err := TestClock(timestampLayout, tc.timestamp)
		if err != nil {
			t.Fatalf("[%s] got error: %v", tc.name, err)
		}
		var c BrokenClock
		err = json.Unmarshal([]byte(tc.jsonStr), &c)
		if err == nil && tc.err {
			t.Fatalf("[%s] expected an error, but didn't get one; got %v instead", tc.name, c)
		} else if err != nil && !tc.err {
			t.Fatalf("[%s] got error: %v", tc.name, err)
		} else if err != nil {
			t.Logf("[%s] got expected error: %v", tc.name, err)
		} else if err == nil && !c.Equal(expected) {
			t.Fatalf("[%s] expected %v; got %v", tc.name, expected, c)
		}
	}
}
