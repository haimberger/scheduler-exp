package task

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/haimberger/scheduler/clock"
)

// Task contains information about something that needs to be done.
type Task struct {
	// Title is a short description of the task.
	Title string `json:"title"`
	// Link is a URL pointing to a Trello ticket, email or similar.
	Link string `json:"link"`
	// Assigner is the name of person who assigned the task.
	Assigner string `json:"assigner"`
	// CanPreempt is true iff this task can preempt long-running higher-priority tasks.
	CanPreempt bool `json:"canPreempt"`
	// IsCancelled is true iff the task is no longer relevant.
	IsCancelled bool `json:"isCancelled"`
	// IsCompleted is true iff the task was successfully completed.
	IsCompleted bool `json:"isCompleted"`
	// Priority represents the task's importance relative to other tasks (low number -> high priority).
	Priority int `json:"priority"`
	// Duration is an estimate of how long the task will take in hours.
	Duration float32 `json:"duration"`
	// CreationTime is when the task was created.
	CreationTime time.Time `json:"creationTime"`
	// StartTime is when the task should jump to the top priority.
	StartTime time.Time `json:"startTime"`
	// ActiveIntervals are times when the task was actively being worked on.
	ActiveIntervals []Interval `json:"activeIntervals"`
	// Clock keeps time.
	Clock clock.Clock `json:"clock"`
}

// Interval specifies a time interval.
type Interval struct {
	// Begin is when the interval starts.
	Begin time.Time `json:"begin"`
	// End is when the interval ends.
	End time.Time `json:"end"`
}

// MkTask creates a new task based on the specified one.
func MkTask(t Task) (Task, error) {
	// TODO: check values (e.g. is title non-empty? is duration greater than zero?)

	// determine clock that will be used to keep time
	var c clock.Clock
	if t.Clock == nil {
		c = clock.StandardClock{}
	} else {
		c = t.Clock
	}

	return Task{
		Title:        t.Title,
		Link:         t.Link,
		Assigner:     t.Assigner,
		CanPreempt:   t.CanPreempt,
		Duration:     t.Duration,
		CreationTime: c.Now(),
		StartTime:    t.StartTime,
		Clock:        c,
	}, nil
}

// Update overwrites a task's fields with the specified values.
func (t *Task) Update(src map[string]string) error {
	dst := reflect.ValueOf(t).Elem()
	for fieldName, value := range src {
		f := dst.FieldByName(fieldName)
		if !f.CanSet() {
			return fmt.Errorf("can't set field %s", fieldName)
		}
		switch f.Type().Name() {
		case "string":
			f.SetString(value)
		case "int":
			i, err := strconv.Atoi(value)
			if err != nil {
				return fmt.Errorf(`invalid value "%s" for %s; expected an integer`, value, fieldName)
			}
			f.SetInt(int64(i))
		case "bool":
			switch value {
			case "true":
				f.SetBool(true)
			case "false":
				f.SetBool(false)
			default:
				return fmt.Errorf(`invalid value "%s" for %s; expected "true" or "false"`, value, fieldName)
			}
		default:
			return fmt.Errorf("unsupported type %s", f.Type())
		}
	}
	return nil
}
