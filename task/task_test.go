package task

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/haimberger/scheduler/clock"
)

var update = flag.Bool("update", false, "update .golden files")

func TestMkTask(t *testing.T) {
	key := "MkTask"
	tc, err := getTestCase(key)
	if err != nil {
		t.Fatal(err)
	}

	// try using default (standard) clock first
	before := time.Now()
	actualTask, err := MkTask(tc)
	if err != nil {
		t.Fatal(err)
	}
	after := time.Now()
	if !before.Before(after) {
		t.Fatal("time is going backwards?!")
	} else if !before.Before(actualTask.CreationTime) {
		t.Fatalf("expected %v to be before %v; it wasn't", before, actualTask.CreationTime)
	} else if !after.After(actualTask.CreationTime) {
		t.Fatalf("expected %v to be after %v; it wasn't", after, actualTask.CreationTime)
	}

	// use broken clock so that tasks created at different times can be compared
	tc.Clock = clock.BrokenClock{}
	actualTask, err = MkTask(tc)
	if err != nil {
		t.Fatal(err)
	}
	actual, err := json.Marshal(actualTask)
	if err != nil {
		t.Fatal(err)
	}
	if err = checkResult(key, actual); err != nil {
		t.Fatal(err)
	}
}

func getTestCase(key string) (Task, error) {
	var t Task

	// read input file
	input, err := ioutil.ReadFile(fmt.Sprintf("testdata/%s.input", key))
	if err != nil {
		return t, err
	}

	// unmarshal the contents of the input file
	json.Unmarshal(input, &t)

	return t, nil
}

func checkResult(key string, actual []byte) error {
	// read golden file, which contains the expected output
	out := fmt.Sprintf("testdata/%s.golden", key)
	expected, err := ioutil.ReadFile(out)
	if err != nil {
		return err
	}

	// check if the actual result matches the contents of the golden file
	if !bytes.Equal(actual, expected) {
		if *update {
			// overwrite the golden file with the actual result
			if err = ioutil.WriteFile(out, actual, 0666); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("expected %s, got %s", expected, actual)
		}
	}
	return nil
}
