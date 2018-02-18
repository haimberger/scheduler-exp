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
	if err := test.LoadInput("basic.input", &cf); err != nil {
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
	if err = test.CompareResults(&task, "basic.golden"); err != nil {
		t.Fatal(err)
	}
}
