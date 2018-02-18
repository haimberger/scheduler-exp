package task

import (
	"testing"
	"time"

	"github.com/haimberger/scheduler/clock"
	"github.com/haimberger/scheduler/test"
)

func TestMkTask(t *testing.T) {
	// load task from file
	var task Task
	if err := test.LoadInput("basic.input", &task); err != nil {
		t.Fatal(err)
	}

	// try using default (standard) clock first
	before := time.Now()
	actual, err := MkTask(task)
	if err != nil {
		t.Fatal(err)
	}
	after := time.Now()
	if !before.Before(after) {
		t.Fatal("time is going backwards?!")
	} else if !before.Before(actual.CreationTime) {
		t.Fatalf("expected %v to be before %v; it wasn't", before, actual.CreationTime)
	} else if !after.After(actual.CreationTime) {
		t.Fatalf("expected %v to be after %v; it wasn't", after, actual.CreationTime)
	}

	// use broken clock so that tasks created at different times can be compared
	task.Clock = &clock.BrokenClock{}
	actual, err = MkTask(task)
	if err != nil {
		t.Fatal(err)
	}
	if err = test.CompareResults(&actual, "basic.golden"); err != nil {
		t.Fatal(err)
	}
}
