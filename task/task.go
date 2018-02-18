package task

import (
	"errors"
	"time"

	"github.com/haimberger/scheduler/clock"
)

// Task contains information about something that needs to be done.
type Task struct {
	// Config contains configurable properties of the task.
	Config Config `json:"config"`
	// CreationTime is when the task was created.
	CreationTime time.Time `json:"creationTime"`
	// ActiveIntervals are times when the task was actively being worked on.
	ActiveIntervals []Interval `json:"activeIntervals"`
}

// Config contains configurable properties of a task.
type Config struct {
	// Title is a short description of the task.
	Title string `json:"title"`
	// Link is a URL pointing to a Trello ticket, email or similar.
	Link string `json:"link"`
	// Assigner is the name of the person who assigned the task.
	Assigner string `json:"assigner"`
	// Priority represents the task's importance relative to other tasks (low number -> high priority).
	Priority int `json:"priority"`
	// Duration is an estimate of how long the task will take in hours.
	Duration float32 `json:"duration"`
	// CanPreempt is true iff this task can preempt long-running higher-priority tasks.
	CanPreempt bool `json:"canPreempt"`
	// IsCancelled is true iff the task is no longer relevant.
	IsCancelled bool `json:"isCancelled"`
	// IsCompleted is true iff the task was successfully completed.
	IsCompleted bool `json:"isCompleted"`
	// StartTime is when the task should jump to the top priority.
	StartTime time.Time `json:"startTime"`
}

// Interval specifies a time interval.
type Interval struct {
	// Begin is when the interval starts.
	Begin time.Time `json:"begin"`
	// End is when the interval ends.
	End time.Time `json:"end"`
}

// MkTask creates a new task based on the specified one.
func MkTask(cf Config, c clock.Clock) (Task, error) {
	// TODO: check values (e.g. is title non-empty? is duration greater than zero?)

	// use standard clock by default
	if c == nil {
		c = &clock.StandardClock{}
	}

	return Task{
		Config: Config{
			Title:      cf.Title,
			Link:       cf.Link,
			Assigner:   cf.Assigner,
			CanPreempt: cf.CanPreempt,
			Duration:   cf.Duration,
			StartTime:  cf.StartTime,
		},
		CreationTime: c.Now(),
	}, nil
}

// Start marks a task as being actively worked on.
func (t *Task) Start(c clock.Clock) error {
	// check if it even makes sense for the task
	if t.Config.IsCancelled {
		return errors.New("can't start cancelled task")
	}
	if t.Config.IsCompleted {
		return errors.New("can't start completed task")
	}

	// check if the task is already marked as being in progress
	if l := len(t.ActiveIntervals); l == 0 || !t.ActiveIntervals[l-1].End.IsZero() {
		// if not, record the beginning of the current interval of work
		t.ActiveIntervals = append(t.ActiveIntervals, Interval{
			Begin: c.Now(),
		})
	}

	return nil
}

// Stop marks a task as no longer being actively worked on.
func (t *Task) Stop(c clock.Clock) error {
	// check if the task is currently marked as being in progress
	if l := len(t.ActiveIntervals); l > 0 && t.ActiveIntervals[l-1].End.IsZero() {
		// if so, record the end of the last interval of work
		t.ActiveIntervals[l-1].End = c.Now()
	}
	return nil
}

// Update overwrites a task's fields with the specified values.
func (t *Task) Update(cf Config) error {
	// TODO: check values (e.g. is title non-empty? is duration greater than zero?)
	// TODO: if isCancelled or isCompleted is set, update active intervals accordingly
	t.Config = cf
	return nil
}
