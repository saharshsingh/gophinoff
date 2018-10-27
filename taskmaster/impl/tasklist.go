package impl

import (
	"github.com/saharshsingh/gophinoff/datastructures"
	"github.com/saharshsingh/gophinoff/datastructures/concurrent"
	"github.com/saharshsingh/gophinoff/datastructures/linkedlist"
)

// CreatePrioritizedTaskList creates a new TaskMaster with no scheduled tasks
func CreatePrioritizedTaskList() *PrioritizedTaskList {
	return &PrioritizedTaskList{concurrent.WrapWithConcurrentList(linkedlist.CreateSorted())}
}

// PrioritizedTaskList is used to schedule and retrieve tasks in priority order
type PrioritizedTaskList struct {
	taskList datastructures.LinkedList
}

// AddNew a new task based on the specified details to backlog
func (taskList *PrioritizedTaskList) AddNew(name string, description string, priority int) *Task {
	task := CreateTask(name, description, priority)
	taskList.Add(task)
	return task
}

// Add the specified task to backlog
func (taskList *PrioritizedTaskList) Add(task *Task) {
	taskList.taskList.Add(task)
}

// PendingCount is the number of tasks still scheduled
func (taskList *PrioritizedTaskList) PendingCount() uint64 {
	return taskList.taskList.Size()
}

// PeekNext will retreive the highest priority task without deleting it
func (taskList *PrioritizedTaskList) PeekNext() *Task {
	return resolveNilOrTask(taskList.taskList.Get(0, true))
}

// Next will delete and return the highest priority task, essentially marking it done
func (taskList *PrioritizedTaskList) Next() *Task {
	return resolveNilOrTask(taskList.taskList.Delete(0, true))
}

func resolveNilOrTask(nilOrTask interface{}) *Task {
	if nilOrTask == nil {
		return nil
	}
	return nilOrTask.(*Task)
}
