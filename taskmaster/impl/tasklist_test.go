package impl

import (
	"testing"
	"time"

	"github.com/saharshsingh/gophinoff/testutils"
)

func TestTypicalFlow(t *testing.T) {

	taskList := CreatePrioritizedTaskList()

	taskList.AddNew("low 1", "first low inserted", 25)
	time.Sleep(1 * time.Millisecond)
	taskList.AddNew("med 1", "first med inserted", 50)
	time.Sleep(1 * time.Millisecond)
	taskList.AddNew("high 1", "first high inserted", 100)
	time.Sleep(1 * time.Millisecond)
	taskList.AddNew("med 2", "second med inserted", 50)
	time.Sleep(1 * time.Millisecond)
	taskList.AddNew("med 3", "third med inserted", 50)
	time.Sleep(1 * time.Millisecond)
	taskList.AddNew("low 2", "second low inserted", 25)
	time.Sleep(1 * time.Millisecond)
	taskList.AddNew("high 2", "second high inserted", 100)
	time.Sleep(1 * time.Millisecond)
	taskList.AddNew("high 3", "third high inserted", 100)
	time.Sleep(1 * time.Millisecond)
	taskList.AddNew("low 3", "third low inserted", 25)

	testutils.AssertEquals("Expected nine scheduled tasks", uint64(9), taskList.PendingCount(), t)

	testutils.AssertEquals("Expected 'high 1'", "high 1", taskList.PeekNext().Name, t)
	testutils.AssertEquals("Expected nine scheduled tasks", uint64(9), taskList.PendingCount(), t)
	testutils.AssertEquals("Expected 'high 1'", "high 1", taskList.Next().Name, t)
	testutils.AssertEquals("Expected eight scheduled tasks", uint64(8), taskList.PendingCount(), t)

	testutils.AssertEquals("Expected 'high 2'", "high 2", taskList.PeekNext().Name, t)
	testutils.AssertEquals("Expected eight scheduled tasks", uint64(8), taskList.PendingCount(), t)
	testutils.AssertEquals("Expected 'high 2'", "high 2", taskList.Next().Name, t)
	testutils.AssertEquals("Expected seven scheduled tasks", uint64(7), taskList.PendingCount(), t)

	testutils.AssertEquals("Expected 'high 3'", "high 3", taskList.PeekNext().Name, t)
	testutils.AssertEquals("Expected seven scheduled tasks", uint64(7), taskList.PendingCount(), t)
	testutils.AssertEquals("Expected 'high 3'", "high 3", taskList.Next().Name, t)
	testutils.AssertEquals("Expected six scheduled tasks", uint64(6), taskList.PendingCount(), t)

	testutils.AssertEquals("Expected 'med 1'", "med 1", taskList.PeekNext().Name, t)
	testutils.AssertEquals("Expected 'med 1'", "med 1", taskList.Next().Name, t)

	testutils.AssertEquals("Expected 'med 2'", "med 2", taskList.PeekNext().Name, t)
	testutils.AssertEquals("Expected 'med 2'", "med 2", taskList.Next().Name, t)

	testutils.AssertEquals("Expected 'med 3'", "med 3", taskList.PeekNext().Name, t)
	testutils.AssertEquals("Expected 'med 3'", "med 3", taskList.Next().Name, t)

	testutils.AssertEquals("Expected 'low 1'", "low 1", taskList.PeekNext().Name, t)
	testutils.AssertEquals("Expected 'low 1'", "low 1", taskList.Next().Name, t)

	testutils.AssertEquals("Expected 'low 2'", "low 2", taskList.PeekNext().Name, t)
	testutils.AssertEquals("Expected 'low 2'", "low 2", taskList.Next().Name, t)

	testutils.AssertEquals("Expected 'low 3'", "low 3", taskList.PeekNext().Name, t)
	testutils.AssertEquals("Expected 'low 3'", "low 3", taskList.Next().Name, t)

	testutils.AssertEquals("Expected zero scheduled tasks", uint64(0), taskList.PendingCount(), t)
}
