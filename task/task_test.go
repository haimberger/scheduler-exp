package task

import (
	"testing"
	"time"

	"github.com/haimberger/scheduler/clock"
	"github.com/haimberger/scheduler/test"
)

func TestMkTask(t *testing.T) {
	// load task config from file
	var cf Config
	if err := test.LoadInput("config", &cf); err != nil {
		t.Fatal(err)
	}

	// try using default (standard) clock first
	before := time.Now()
	task, err := MkTask(cf, nil)
	if err != nil {
		t.Fatal(err)
	}
	after := time.Now()
	if !before.Before(after) {
		t.Fatal("time is going backwards?!")
	} else if !before.Before(task.CreationTime) {
		t.Fatalf("expected %v to be before %v; it wasn't", before, task.CreationTime)
	} else if !after.After(task.CreationTime) {
		t.Fatalf("expected %v to be after %v; it wasn't", after, task.CreationTime)
	}

	// use broken clock so that tasks created at different times can be compared
	task, err = MkTask(cf, &clock.BrokenClock{})
	if err != nil {
		t.Fatal(err)
	}
	if err = test.CompareResults(&task, "new"); err != nil {
		t.Fatal(err)
	}
}

func TestUpdate(t *testing.T) {
	type testCase struct {
		name string
		in   string
		out  string
	}
	tcs := []testCase{
		{name: "basic", in: "changes", out: "updated"},
	}

	for _, tc := range tcs {
		// load task config from file
		var cf Config
		if err := test.LoadInput("config", &cf); err != nil {
			t.Fatal(err)
		}
		task, err := MkTask(cf, &clock.BrokenClock{})
		if err != nil {
			t.Fatal(err)
		}

		// load changes from file
		var changes Config
		if err := test.LoadInput(tc.in, &changes); err != nil {
			t.Errorf("[%s] got error: %v", tc.name, err)
			return
		}

		if err := task.Update(changes); err != nil {
			t.Errorf("[%s] got error: %v", tc.name, err)
			return
		}

		if err := test.CompareResults(task, tc.out); err != nil {
			t.Errorf("[%s] got error: %v", tc.name, err)
		}
	}
}
