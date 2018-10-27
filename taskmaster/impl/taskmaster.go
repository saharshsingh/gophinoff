package impl

import (
	"fmt"
)

// CreateTaskMaster returns a new task master with no scheduled/assigned tasks
func CreateTaskMaster() *TaskMaster {
	return &TaskMaster{CreatePrioritizedTaskList(), map[string]*Task{}}
}

// TaskMaster is used to prioritize and assign tasks
type TaskMaster struct {
	tasks         *PrioritizedTaskList
	assignedTasks map[string]*Task
}

// Create a new task
func (master *TaskMaster) Create(name string, description string, priority int) {
	master.tasks.AddNew(name, description, priority)
}

// Assign the next available task to specified assignee
func (master *TaskMaster) Assign(assignee string) (*Task, error) {

	if master.assignedTasks[assignee] != nil {
		return nil, fmt.Errorf("'%v' is already assigned another task", assignee)
	}

	task := master.tasks.Next()
	if task == nil {
		return nil, fmt.Errorf("No pending tasks avaialble to assign")
	}

	master.assignedTasks[assignee] = task

	return task, nil
}

// PeekNext returns the next pending task
func (master *TaskMaster) PeekNext() *Task {
	return master.tasks.PeekNext()
}

// PeekAssigned returns pointer to task currently assigned to assignee
func (master *TaskMaster) PeekAssigned(assignee string) *Task {
	return master.assignedTasks[assignee]
}

// Unassign task assigned to specified assignee. Assignee will have nothing assigned, and assigned task is added back to list
func (master *TaskMaster) Unassign(assignee string) error {

	unassigned, error := master.unassign(assignee)

	if unassigned != nil {
		master.tasks.Add(unassigned)
	}

	return error
}

// MarkDone unassigns and deletes task currently assigned to assignee
func (master *TaskMaster) MarkDone(assignee string) error {
	_, error := master.unassign(assignee)
	return error
}

func (master *TaskMaster) unassign(assignee string) (*Task, error) {

	assigned := master.assignedTasks[assignee]

	if assigned == nil {
		return nil, fmt.Errorf("No task assigned to '%v'", assignee)
	}

	delete(master.assignedTasks, assignee)

	return assigned, nil
}
