package linkedlist

import (
	"fmt"

	"github.com/saharshsingh/gophinoff/datastructures"
)

// CreateSorted creates a linked list that guarantees elements are inserted
// in sorted order. Elements must either be a supported type or satisfy the
// Comparable interface to be accepted. Following types are supported out of box:
//
// 	- int
// 	- string
func CreateSorted() datastructures.LinkedList {
	return &sortedLinkedList{Create()}
}

// Add inserts value in the list at correct position based on natural order
func (list *sortedLinkedList) Add(value interface{}) {

	comparable := createComparable(value)

	if list.size == 0 {

		list.LinkedList.Add(comparable.value)

	} else {

		nodeToPrepend := list.head
		setAsNewHead := true

		for comparable.compare(nodeToPrepend.value) >= 0 {
			setAsNewHead = false

			nodeToPrepend = nodeToPrepend.next
			if nodeToPrepend == list.head {
				break
			}
		}

		inserted := nodeToPrepend.prepend(comparable.value)
		if setAsNewHead {
			list.head = inserted
		}

		list.size++
	}

}

// -- internal

type sortedLinkedList struct {
	*LinkedList
}

func createComparable(value interface{}) *comparable {

	_string, ok := value.(string)
	if ok {
		return &comparable{_string}
	}

	_int, ok := value.(int)
	if ok {
		return &comparable{_int}
	}

	_Comparable, ok := value.(datastructures.Comparable)
	if ok {
		return &comparable{_Comparable}
	}

	panic(fmt.Sprintf("Must be int, string, or github.com/saharshsingh/gophinoff/datastructures.Comparable"))
}

type comparable struct {
	value interface{}
}

func (c *comparable) compare(other interface{}) int {

	switch c.value.(type) {

	case string:
		c1 := c.value.(string)
		c2 := other.(string)
		if c1 < c2 {
			return -1
		}
		if c1 == c2 {
			return 0
		}
		return 1

	case int:
		c1 := c.value.(int)
		c2 := other.(int)
		if c1 < c2 {
			return -1
		}
		if c1 == c2 {
			return 0
		}
		return 1

	}

	c1 := c.value.(datastructures.Comparable)
	c2 := other.(datastructures.Comparable)
	return c1.Compare(c2)
}
