package linkedlist

import (
	"github.com/saharshsingh/gophinoff/datastructures"
)

// LinkedList implements a circular doubly linked list
type LinkedList struct {
	head *node
	size uint64
}

// Create an empty LinkedList
func Create() *LinkedList {
	return &LinkedList{nil, 0}
}

// Size of list
func (list *LinkedList) Size() uint64 {
	return list.size
}

// ReversibleIterator is an iterator that can be optionally configured to start from tail and iterate backwards
func (list *LinkedList) ReversibleIterator(flipped bool) datastructures.Iterator {
	head := list.head
	if flipped {
		head = head.prev
	}
	return &linkedListIterator{head, head, flipped}
}

// Iterator is used to iterate over list values
func (list *LinkedList) Iterator() datastructures.Iterator {
	return list.ReversibleIterator(false)
}

// Insert a value at specified offset. Offset must be less than or equal to list size
func (list *LinkedList) Insert(value interface{}, offset uint64, fromBack bool) {

	if offset == uint64(0) && list.size == 0 {

		node := &node{value, nil, nil}
		list.head, node.prev, node.next = node, node, node
		list.size = 1

	} else {

		var nodeToPrepend *node
		if offset == list.size {
			nodeToPrepend = list.head
		} else {
			nodeToPrepend = list.getAtOffset(offset, fromBack)
		}

		setInsertedNodeAsListHead := offset == 0
		if fromBack {
			setInsertedNodeAsListHead = offset == list.size-1
		}

		insertedNode := nodeToPrepend.prepend(value)
		list.size++

		if setInsertedNodeAsListHead {
			list.head = insertedNode
		}

	}
}

// Add adds a new node to the tail of the list
func (list *LinkedList) Add(value interface{}) {
	list.Insert(value, list.size, false)
}

// Clear clears the list to become empty
func (list *LinkedList) Clear() {
	list.head = nil
	list.size = 0
}

// Get value stored at specified index
func (list *LinkedList) Get(offset uint64, fromBack bool) interface{} {
	return list.getAtOffset(offset, fromBack).value
}

// Delete value stored at specified index
func (list *LinkedList) Delete(offset uint64, fromBack bool) interface{} {

	nodeToDelete := list.getAtOffset(offset, fromBack)
	deletedValue := nodeToDelete.value

	// reset list head if current list head is being deleted
	if list.size == 1 {
		list.head = nil
	} else {

		resetHead := offset == 0
		if fromBack {
			resetHead = offset == list.size-1
		}

		if resetHead {
			list.head = list.head.next
		}
	}

	// delete node
	nodeToDelete.delete()
	list.size--

	return deletedValue
}

func (list *LinkedList) getAtOffset(offset uint64, fromBack bool) *node {

	iter := list.ReversibleIterator(fromBack).(*linkedListIterator)

	for i := uint64(0); i < offset; i++ {
		iter.next()
	}

	return iter.next()
}
