package linkedlist

import (
	"testing"

	"github.com/saharshsingh/gophinoff/datastructures"
	"github.com/saharshsingh/gophinoff/testutils"
)

func TestCreate(t *testing.T) {

	list := Create()
	testutils.AssertEquals("Size must be 0 at create", uint64(0), list.Size(), t)

	testutils.AssertFalse("HasNext() should return false for empty list", list.Iterator().HasNext(), t)
}

func TestSatisfiesLinkedListInterface(t *testing.T) {
	var list datastructures.LinkedList = Create()
	testutils.AssertEquals("Expected Size() to be 0", uint64(0), list.Size(), t)
}

func TestIteratorWhenEmpty(t *testing.T) {
	testutils.AssertFalse("Expected iterator to be empty", Create().Iterator().HasNext(), t)
}

func TestIterator(t *testing.T) {

	list := Create()
	list.Add("One")
	list.Add("Two")
	list.Add("Three")
	list.Add("Four")
	testutils.AssertEquals("Expected Size to be 4", uint64(4), list.Size(), t)

	iterator := list.Iterator()
	testutils.AssertTrue("Expected HasNext() to return true", iterator.HasNext(), t)
	testutils.AssertEquals("Expected 'One'", "One", iterator.Next(), t)

	testutils.AssertTrue("Expected HasNext() to return true", iterator.HasNext(), t)
	testutils.AssertEquals("Expected 'Two'", "Two", iterator.Next(), t)

	testutils.AssertTrue("Expected HasNext() to return true", iterator.HasNext(), t)
	testutils.AssertEquals("Expected 'Three'", "Three", iterator.Next(), t)

	testutils.AssertTrue("Expected HasNext() to return true", iterator.HasNext(), t)
	testutils.AssertEquals("Expected 'Four'", "Four", iterator.Next(), t)

	testutils.AssertFalse("HasNext() should return false after traversing all elements", iterator.HasNext(), t)
}

func TestReversibleIteratorWhenEmpty(t *testing.T) {

	// normal iterator
	testutils.AssertFalse("Expected iterator to be empty", Create().ReversibleIterator(false).HasNext(), t)

	// reverse iterator
	testutils.AssertFalse("Expected flipped iterator to be empty", Create().ReversibleIterator(true).HasNext(), t)
}

func TestReversibleIterator(t *testing.T) {

	list := Create()
	list.Add("One")
	list.Add("Two")
	list.Add("Three")
	list.Add("Four")
	testutils.AssertEquals("Expected Size to be 4", uint64(4), list.Size(), t)

	reverseIterator := list.ReversibleIterator(true)
	testutils.AssertTrue("Expected first HasNext() to return true", reverseIterator.HasNext(), t)
	testutils.AssertEquals("Expected 'Four'", "Four", reverseIterator.Next(), t)

	testutils.AssertTrue("Expected second HasNext() to return true", reverseIterator.HasNext(), t)
	testutils.AssertEquals("Expected 'Three'", "Three", reverseIterator.Next(), t)

	testutils.AssertTrue("Expected third HasNext() to return true", reverseIterator.HasNext(), t)
	testutils.AssertEquals("Expected 'Two'", "Two", reverseIterator.Next(), t)

	testutils.AssertTrue("Expected fourth HasNext() to return true", reverseIterator.HasNext(), t)
	testutils.AssertEquals("Expected 'One'", "One", reverseIterator.Next(), t)

	testutils.AssertFalse("HasNext() should return false after traversing all elements", reverseIterator.HasNext(), t)
}

func TestInsert(t *testing.T) {

	list := Create()
	list.Add("Two")
	list.Add("Three")
	list.Add("Six")
	testutils.AssertEquals("Expected Size to be 3", uint64(3), list.Size(), t)

	list.Insert("One", 0, false)
	testutils.AssertEquals("Expected Size to be 4", uint64(4), list.Size(), t)

	list.Insert("Four", 3, false)
	testutils.AssertEquals("Expected Size to be 5", uint64(5), list.Size(), t)

	list.Insert("Five", 4, false)
	testutils.AssertEquals("Expected Size to be 6", uint64(6), list.Size(), t)

	list.Insert("Seven", 6, false)
	testutils.AssertEquals("Expected Size to be 7", uint64(7), list.Size(), t)

	testutils.AssertEquals("Expected 'One'", "One", list.Get(0, false), t)
	testutils.AssertEquals("Expected 'Two'", "Two", list.Get(1, false), t)
	testutils.AssertEquals("Expected 'Three'", "Three", list.Get(2, false), t)
	testutils.AssertEquals("Expected 'Four'", "Four", list.Get(3, false), t)
	testutils.AssertEquals("Expected 'Five'", "Five", list.Get(4, false), t)
	testutils.AssertEquals("Expected 'Six'", "Six", list.Get(5, false), t)
	testutils.AssertEquals("Expected 'Seven'", "Seven", list.Get(6, false), t)

}

func TestInsertFromBack(t *testing.T) {

	list := Create()
	list.Add("Two")
	list.Add("Three")
	list.Add("Six")
	testutils.AssertEquals("Expected Size to be 3", uint64(3), list.Size(), t)

	list.Insert("One", 2, true)
	testutils.AssertEquals("Expected Size to be 4", uint64(4), list.Size(), t)

	list.Insert("Four", 0, true)
	testutils.AssertEquals("Expected Size to be 5", uint64(5), list.Size(), t)

	list.Insert("Five", 0, true)
	testutils.AssertEquals("Expected Size to be 6", uint64(6), list.Size(), t)

	list.Insert("Seven", 6, true)
	testutils.AssertEquals("Expected Size to be 7", uint64(7), list.Size(), t)

	testutils.AssertEquals("Expected 'One'", "One", list.Get(0, false), t)
	testutils.AssertEquals("Expected 'Two'", "Two", list.Get(1, false), t)
	testutils.AssertEquals("Expected 'Three'", "Three", list.Get(2, false), t)
	testutils.AssertEquals("Expected 'Four'", "Four", list.Get(3, false), t)
	testutils.AssertEquals("Expected 'Five'", "Five", list.Get(4, false), t)
	testutils.AssertEquals("Expected 'Six'", "Six", list.Get(5, false), t)
	testutils.AssertEquals("Expected 'Seven'", "Seven", list.Get(6, false), t)

}

func TestClearWhenFull(t *testing.T) {

	list := Create()
	list.Add("One")
	list.Add("Two")
	list.Add("Three")
	list.Add("Four")

	list.Clear()

	testutils.AssertEquals("Expected Size to be 0", uint64(0), list.Size(), t)
	testutils.AssertFalse("Expected HasNext() to return false", list.Iterator().HasNext(), t)
}

func TestClearWhenEmpty(t *testing.T) {

	list := Create()

	list.Clear()

	testutils.AssertEquals("Expected Size to be 0", uint64(0), list.Size(), t)
	testutils.AssertFalse("Expected HasNext() to return false", list.Iterator().HasNext(), t)
}

func TestGetWhenMany(t *testing.T) {

	list := Create()
	list.Add("One")
	list.Add("Two")
	list.Add("Three")
	list.Add("Four")

	testutils.AssertEquals("Expected 'Two'", "Two", list.Get(uint64(1), false), t)
	testutils.AssertEquals("Expected 'Four'", "Four", list.Get(uint64(3), false), t)

	testutils.AssertEquals("Expected 'Two'", "Two", list.Get(uint64(2), true), t)
	testutils.AssertEquals("Expected 'Four'", "Four", list.Get(uint64(0), true), t)
}

func TestGetWhenOne(t *testing.T) {

	list := Create()
	list.Add("One")

	testutils.AssertEquals("Expected 'One'", "One", list.Get(uint64(0), false), t)
	testutils.AssertEquals("Expected 'One'", "One", list.Get(uint64(0), true), t)
}

func TestGetWhenOffsetOutOfBoundsFromFront(t *testing.T) {

	list := Create()
	list.Add("One")
	list.Add("Two")

	defer testutils.AssertPanic("Expected out of bounds 'Get()' to cause panic", t)

	list.Get(2, false)
}

func TestGetWhenOffsetOutOfBoundsFromBack(t *testing.T) {

	list := Create()
	list.Add("One")
	list.Add("Two")

	defer testutils.AssertPanic("Expected out of bounds 'Get()' to cause panic", t)

	list.Get(2, true)
}

func TestDeleteWhenMany(t *testing.T) {

	list := Create()
	list.Add("One")
	list.Add("Two")
	list.Add("Three")
	list.Add("Four")

	testutils.AssertEquals("Expected 'Two'", "Two", list.Delete(1, false), t)
	testutils.AssertEquals("Expected 'One'", "One", list.Delete(2, true), t)

	testutils.AssertEquals("Expected Size to be 2", uint64(2), list.Size(), t)

	iterator := list.Iterator()

	testutils.AssertTrue("Expected HasNext() to return true", iterator.HasNext(), t)
	testutils.AssertEquals("Expected 'Three'", "Three", iterator.Next(), t)

	testutils.AssertTrue("Expected HasNext() to return true", iterator.HasNext(), t)
	testutils.AssertEquals("Expected 'Four'", "Four", iterator.Next(), t)

	testutils.AssertFalse("Expected HasNext() to return false", iterator.HasNext(), t)
}

func TestDeleteWhenOne(t *testing.T) {

	list := Create()

	list.Add("One")
	testutils.AssertEquals("Expected 'One'", "One", list.Delete(0, false), t)
	testutils.AssertEquals("Expected Size to be 0", uint64(0), list.Size(), t)
	testutils.AssertFalse("Expected HasNext() to return false", list.Iterator().HasNext(), t)

	list.Add("One")
	testutils.AssertEquals("Expected 'One'", "One", list.Delete(0, true), t)
	testutils.AssertEquals("Expected Size to be 0", uint64(0), list.Size(), t)
	testutils.AssertFalse("Expected HasNext() to return false", list.Iterator().HasNext(), t)
}

func TestDeleteWhenOffsetOutOfBoundsFromFront(t *testing.T) {

	list := Create()
	list.Add("One")
	list.Add("Two")

	defer testutils.AssertPanic("Expected out of bounds 'Get()' to cause panic", t)

	list.Delete(2, false)
}

func TestDeleteWhenOffsetOutOfBoundsFromBack(t *testing.T) {

	list := Create()
	list.Add("One")
	list.Add("Two")

	defer testutils.AssertPanic("Expected out of bounds 'Get()' to cause panic", t)

	list.Delete(2, true)
}
