package task

import (
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
	// Clock keeps time.
	Clock clock.Clock `json:"clock"`
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
		Clock:        c,
	}, nil
}

// Update overwrites a task's fields with the specified values.
func (t *Task) Update(cf Config) error {
	// TODO: check values (e.g. is title non-empty? is duration greater than zero?)
	t.Config = cf
	return nil
}
