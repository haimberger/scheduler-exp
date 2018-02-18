package task

import (
	"testing"
	"time"

	"github.com/haimberger/scheduler/clock"
	"github.com/haimberger/scheduler/test"
)

func TestMkTask_standard(t *testing.T) {
	var cf Config
	if err := test.LoadInput("config", &cf); err != nil {
		t.Fatal(err)
	}

	before := time.Now()
	task, err := MkTask(cf, nil)
	if err != nil {
		t.Fatal(err)
	}
	after := time.Now()

	if !before.Before(after) {
		t.Error("time is going backwards?!")
	} else if !before.Before(task.CreationTime) {
		t.Errorf("expected %v to be before %v; it wasn't", before, task.CreationTime)
	} else if !after.After(task.CreationTime) {
		t.Errorf("expected %v to be after %v; it wasn't", after, task.CreationTime)
	}
}

func TestMkTask_broken(t *testing.T) {
	var cf Config
	if err := test.LoadInput("config", &cf); err != nil {
		t.Fatal(err)
	}
	task, err := MkTask(cf, &clock.BrokenClock{})
	if err != nil {
		t.Fatal(err)
	}
	if err = test.CompareResults(&task, "new"); err != nil {
		t.Error(err)
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
		if err := testUpdate(tc.in, tc.out); err != nil {
			t.Errorf("[%s] %v", tc.name, err)
		}
	}
}

func testUpdate(in, out string) error {
	var cf Config
	if err := test.LoadInput("config", &cf); err != nil {
		return err
	}
	task, err := MkTask(cf, &clock.BrokenClock{})
	if err != nil {
		return err
	}
	var changes Config
	if err := test.LoadInput(in, &changes); err != nil {
		return err
	}
	if err := task.Update(changes); err != nil {
		return err
	}
	return test.CompareResults(task, out)
}
