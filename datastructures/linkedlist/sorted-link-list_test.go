package linkedlist

import (
	"testing"

	"github.com/saharshsingh/gophinoff/datastructures"
	"github.com/saharshsingh/gophinoff/testutils"
)

type person struct {
	name string
	age  int
}

func (p *person) Compare(other datastructures.Comparable) int {

	p2 := other.(*person)

	if p.name < p2.name {
		return -1
	}
	if p.name == p2.name {
		return 0
	}
	return 1
}

func TestAddString(t *testing.T) {

	list := CreateSorted()

	list.Add("Joe")
	list.Add("Candace")
	list.Add("Zoe")
	list.Add("Adam")
	list.Add("Bob")
	list.Add("Joe")

	testutils.AssertEquals("Expected Size() to be 6", uint64(6), list.Size(), t)

	iterator := list.ReversibleIterator(true)
	testutils.AssertTrue("Expected HasNext() to return 'true'", iterator.HasNext(), t)
	testutils.AssertEquals("Expected 'Zoe'", "Zoe", iterator.Next(), t)
	testutils.AssertTrue("Expected HasNext() to return 'true'", iterator.HasNext(), t)
	testutils.AssertEquals("Expected 'Joe'", "Joe", iterator.Next(), t)
	testutils.AssertTrue("Expected HasNext() to return 'true'", iterator.HasNext(), t)
	testutils.AssertEquals("Expected 'Joe'", "Joe", iterator.Next(), t)
	testutils.AssertTrue("Expected HasNext() to return 'true'", iterator.HasNext(), t)
	testutils.AssertEquals("Expected 'Candace'", "Candace", iterator.Next(), t)
	testutils.AssertTrue("Expected HasNext() to return 'true'", iterator.HasNext(), t)
	testutils.AssertEquals("Expected 'Bob'", "Bob", iterator.Next(), t)
	testutils.AssertTrue("Expected HasNext() to return 'true'", iterator.HasNext(), t)
	testutils.AssertEquals("Expected 'Adam'", "Adam", iterator.Next(), t)
	testutils.AssertFalse("Expected HasNext() to return 'false'", iterator.HasNext(), t)
}

func TestAddInt(t *testing.T) {

	list := CreateSorted()

	list.Add(4)
	list.Add(3)
	list.Add(5)
	list.Add(1)
	list.Add(2)
	list.Add(4)

	testutils.AssertEquals("Expected Size() to be 6", uint64(6), list.Size(), t)
}

func TestAddComparable(t *testing.T) {

	list := CreateSorted()

	list.Add(&person{"Joe", 13})
	list.Add(&person{"Candace", 14})
	list.Add(&person{"Zoe", 15})
	list.Add(&person{"Adam", 16})
	list.Add(&person{"Bob", 17})
	list.Add(&person{"Joe", 18})

	testutils.AssertEquals("Expected Size() to be 6", uint64(6), list.Size(), t)

	testutils.AssertEquals("Expected 'Adam'", "Adam", list.Get(uint64(0), false).(*person).name, t)
	testutils.AssertEquals("Expected 'Bob'", "Bob", list.Get(uint64(1), false).(*person).name, t)
	testutils.AssertEquals("Expected 'Candace'", "Candace", list.Get(uint64(2), false).(*person).name, t)
	testutils.AssertEquals("Expected 'Joe'", "Joe", list.Get(uint64(3), false).(*person).name, t)
	testutils.AssertEquals("Expected first 'Joe' to be 13", 13, list.Get(uint64(3), false).(*person).age, t)
	testutils.AssertEquals("Expected 'Joe'", "Joe", list.Get(uint64(4), false).(*person).name, t)
	testutils.AssertEquals("Expected second 'Joe' to be 18", 18, list.Get(uint64(4), false).(*person).age, t)
	testutils.AssertEquals("Expected 'Zoe'", "Zoe", list.Get(uint64(5), false).(*person).name, t)
}

func TestAddNonComparable(t *testing.T) {

	list := CreateSorted()

	defer testutils.AssertPanic("Expected next statement to cause panic", t)
	list.Add(&struct{ name string }{"Blah"})
}

func TestAddMixedTypes(t *testing.T) {

	list := CreateSorted()

	list.Add(4)

	defer testutils.AssertPanic("Expected next statement to cause panic", t)
	list.Add("Blah")
}
