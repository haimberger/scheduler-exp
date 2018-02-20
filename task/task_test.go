package task

import (
	"errors"
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

func TestStart(t *testing.T) {
	type testCase struct {
		name      string
		in        string
		out       string
		expectErr bool
	}
	tcs := []testCase{
		{name: "new", in: "new", out: "started"},
		{name: "wip", in: "wip", out: "wip"},
		{name: "paused", in: "paused", out: "restarted"},
		{name: "cancelled", in: "cancelled", expectErr: true},
		{name: "completed", in: "completed", expectErr: true},
	}

	for _, tc := range tcs {
		if err := testStart(tc.in, tc.out, tc.expectErr); err != nil {
			t.Errorf("[%s] %v", tc.name, err)
		}
	}
}

func testStart(in, out string, expectErr bool) error {
	var task Task
	if err := test.LoadInput(in, &task); err != nil {
		return err
	}
	c, err := clock.TestClock(clock.TimestampLayout, "2018-02-16T18:21:34Z")
	if err != nil {
		return err
	}
	err = task.Start(c)
	if expectErr {
		if err == nil {
			return errors.New("expected an error, but didn't get one")
		}
		return nil
	}
	if err != nil {
		return err
	}
	return test.CompareResults(task, out)
}

func TestStop(t *testing.T) {
	type testCase struct {
		name string
		in   string
		out  string
	}
	tcs := []testCase{
		{name: "new", in: "new", out: "new"},
		{name: "wip", in: "wip", out: "paused"},
		{name: "paused", in: "paused", out: "paused"},
		{name: "cancelled", in: "cancelled", out: "cancelled"},
		{name: "completed", in: "completed", out: "completed"},
	}

	for _, tc := range tcs {
		if err := testStop(tc.in, tc.out); err != nil {
			t.Errorf("[%s] %v", tc.name, err)
		}
	}
}

func testStop(in, out string) error {
	var task Task
	if err := test.LoadInput(in, &task); err != nil {
		return err
	}
	c, err := clock.TestClock(clock.TimestampLayout, "2018-02-16T18:21:34Z")
	if err != nil {
		return err
	}
	err = task.Stop(c)
	if err != nil {
		return err
	}
	return test.CompareResults(task, out)
}
