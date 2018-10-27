package impl

import (
	"testing"
	"time"

	"github.com/saharshsingh/gophinoff/datastructures"
	"github.com/saharshsingh/gophinoff/testutils"
)

func TestTaskSatisfiesComparable(t *testing.T) {
	task1 := CreateTask("one", "one", 25)
	task2 := CreateTask("two", "two", 50)

	var comparable datastructures.Comparable = task1
	testutils.AssertEquals("Expected comparable to be less than task2", -1, comparable.Compare(task2), t)
}

func TestPriority(t *testing.T) {
	task1 := CreateTask("one", "one", 25)
	task2 := CreateTask("two", "two", 50)

	testutils.AssertEquals("Expected Priority() to equal 25", 25, task1.Priority(), t)
	testutils.AssertEquals("Expected Priority() to equal 50", 50, task2.Priority(), t)
}

func TestCreatedAt(t *testing.T) {

	startedAt := time.Now()

	time.Sleep(10 * time.Millisecond)
	task1 := CreateTask("one", "one", 25)
	time.Sleep(10 * time.Millisecond)

	endedAt := time.Now()

	testutils.AssertTrue("Expected startedAt timestamp to be before task1 creation timestamp", startedAt.Before(task1.CreatedAt()), t)
	testutils.AssertTrue("Expected endedAt timestamp to be before task1 creation timestamp", endedAt.After(task1.CreatedAt()), t)
}

func TestCompare(t *testing.T) {

	task1 := CreateTask("one", "one", 25)
	time.Sleep(5 * time.Millisecond)
	task2 := CreateTask("two", "two", 50)
	time.Sleep(5 * time.Millisecond)
	task3 := CreateTask("three", "three", 25)

	testutils.AssertEquals("Expected one to equal one", 0, task1.Compare(task1), t)
	testutils.AssertEquals("Expected two to be greater than one", 1, task2.Compare(task1), t)
	testutils.AssertEquals("Expected three to be less than two", -1, task3.Compare(task2), t)
	testutils.AssertEquals("Expected one to be greater than three", 1, task1.Compare(task3), t)
	testutils.AssertEquals("Expected three to be less than one", -1, task3.Compare(task1), t)
}
