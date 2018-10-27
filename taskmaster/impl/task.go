package impl

import (
	"time"

	"github.com/saharshsingh/gophinoff/datastructures"
)

// CreateTask creates a new Task of specified priority
func CreateTask(name string, description string, priority int) *Task {
	return &Task{name, description, priority, time.Now()}
}

// Task represents a task with configurable priority
type Task struct {
	Name        string
	Description string
	priority    int
	createdAt   time.Time
}

// Priority value assigned to task at creation
func (task *Task) Priority() int {
	return task.priority
}

// CreatedAt is the time the task was created
func (task *Task) CreatedAt() time.Time {
	return task.createdAt
}

// Compare implementation to satisfy Comparable interface needed for sort at insertion into list
func (task *Task) Compare(other datastructures.Comparable) int {

	otherTask := other.(*Task)

	priority, oPriority := task.priority, otherTask.priority

	if priority == oPriority {

		createdAt, oCreatedAt := task.createdAt.UnixNano(), otherTask.createdAt.UnixNano()

		if createdAt == oCreatedAt {
			return 0
		}

		if createdAt > oCreatedAt {
			return -1
		}

		return 1
	}

	if priority > oPriority {
		return 1
	}

	return -1
}
