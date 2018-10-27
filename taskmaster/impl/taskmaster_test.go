package impl

import (
	"testing"
	"time"

	"github.com/saharshsingh/gophinoff/testutils"
)

func TestCreateTaskMaster(t *testing.T) {

	taskMaster := CreateTaskMaster()

	var nilTask *Task
	testutils.AssertEquals("Expected PeekNext() to return nil", nilTask, taskMaster.PeekNext(), t)
	testutils.AssertEquals("Expected PeekAssigned() to return nil", nilTask, taskMaster.PeekAssigned("saharsh"), t)

	assignedTask, assignmentError := taskMaster.Assign("saharsh")
	testutils.AssertEquals("Expected error on Assign()", nilTask, assignedTask, t)
	testutils.AssertEquals("Expected error on Assign()", "No pending tasks avaialble to assign", assignmentError.Error(), t)
}

func TestCreate(t *testing.T) {

	taskMaster := CreateTaskMaster()

	taskMaster.Create("Sample1", "Normal Priority Task", 25)
	taskMaster.Create("Sample2", "High Priority Task", 50)
	taskMaster.Create("Sample3", "Normal Priority Task", 25)

	nextTask := taskMaster.PeekNext()
	testutils.AssertEquals("Expected highest priority task at top", "Sample2", nextTask.Name, t)
}

func TestAssign(t *testing.T) {

	taskMaster := createAndInitTaskMaster()

	jimTask, jimError := taskMaster.Assign("jim")
	testutils.AssertEquals("Expected 'Crt1' to be assigned to 'jim'", "Crt1", jimTask.Name, t)
	testutils.AssertEquals("Expected no error in assigning task to 'jim'", nil, jimError, t)

	aliceTask, aliceError := taskMaster.Assign("alice")
	testutils.AssertEquals("Expected 'Crt2' to be assigned to 'alice'", "Crt2", aliceTask.Name, t)
	testutils.AssertEquals("Expected no error in assigning task to 'alice'", nil, aliceError, t)

	jimTask2, jimError2 := taskMaster.Assign("jim")
	testutils.AssertEquals("Expected error in assigning second task to 'jim'", ((*Task)(nil)), jimTask2, t)
	testutils.AssertEquals("Expected error in assigning second task to 'jim'", "'jim' is already assigned another task", jimError2.Error(), t)

	testutils.AssertEquals("Expected next pending task to be 'Crt3'", "Crt3", taskMaster.PeekNext().Name, t)
}

func TestUnassignWhenNoTasksAssigned(t *testing.T) {

	taskMaster := createAndInitTaskMaster()

	testutils.AssertEquals("Expected error when unassigning before assigning", "No task assigned to 'jim'", taskMaster.Unassign("jim").Error(), t)

}

func TestUnassign(t *testing.T) {

	taskMaster := createAndInitTaskMaster()

	taskMaster.Assign("jim")
	testutils.AssertEquals("Expected next pending task to be 'Crt2'", "Crt2", taskMaster.PeekNext().Name, t)

	testutils.AssertEquals("Expected no error unassigning from 'jim'", nil, taskMaster.Unassign("jim"), t)
	testutils.AssertEquals("Expected no task assigned to 'jim'", ((*Task)(nil)), taskMaster.PeekAssigned("jim"), t)
	testutils.AssertEquals("Expected next pending task to be 'Crt1'", "Crt1", taskMaster.PeekNext().Name, t)

}

func TestMarkDoneWhenNoTasksAssigned(t *testing.T) {

	taskMaster := createAndInitTaskMaster()

	testutils.AssertEquals("Expected error when unassigning before assigning", "No task assigned to 'jim'", taskMaster.MarkDone("jim").Error(), t)

}

func TestMarkDone(t *testing.T) {

	taskMaster := createAndInitTaskMaster()

	taskMaster.Assign("jim")
	testutils.AssertEquals("Expected next pending task to be 'Crt2'", "Crt2", taskMaster.PeekNext().Name, t)

	testutils.AssertEquals("Expected no error marking task done for 'jim'", nil, taskMaster.MarkDone("jim"), t)
	testutils.AssertEquals("Expected no task assigned to 'jim'", ((*Task)(nil)), taskMaster.PeekAssigned("jim"), t)
	testutils.AssertEquals("Expected next pending task to be 'Crt2'", "Crt2", taskMaster.PeekNext().Name, t)

}

func createAndInitTaskMaster() *TaskMaster {

	taskMaster := CreateTaskMaster()

	taskMaster.Create("Low1", "Low Priority Task", 25)
	taskMaster.Create("Rtn1", "Routine Priority Task", 50)
	taskMaster.Create("Crt1", "Critical Priority Task", 75)

	time.Sleep(1 * time.Millisecond)
	taskMaster.Create("Low2", "Low Priority Task", 25)
	taskMaster.Create("Rtn2", "Routine Priority Task", 50)
	taskMaster.Create("Crt2", "Critical Priority Task", 75)

	time.Sleep(1 * time.Millisecond)
	taskMaster.Create("Low3", "Low Priority Task", 25)
	taskMaster.Create("Rtn3", "Routine Priority Task", 50)
	taskMaster.Create("Crt3", "Critical Priority Task", 75)

	return taskMaster
}
