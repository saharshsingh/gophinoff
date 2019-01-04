package concurrent

import (
	"sync"

	"github.com/saharshsingh/gophinoff/datastructures"
	"github.com/saharshsingh/gophinoff/datastructures/linkedlist"
)

// LinkedList is a thread safe implementation of local.datastructures.LinkedList
type LinkedList struct {
	innerList       datastructures.LinkedList
	lock            *sync.RWMutex
	numberOfUpdates uint64
}

// CreateConcurrentList creates a thread safe linked list
func CreateConcurrentList() *LinkedList {
	return WrapWithConcurrentList(linkedlist.Create())
}

// WrapWithConcurrentList wraps a thread safe list implementation around provided list
func WrapWithConcurrentList(list datastructures.LinkedList) *LinkedList {
	return &LinkedList{list, &sync.RWMutex{}, 0}
}

// Size of list
func (list *LinkedList) Size() uint64 {

	list.lock.RLock()
	defer list.lock.RUnlock()

	return list.innerList.Size()
}

// ReversibleIterator is an iterator that can be optionally configured to start from tail and iterate backwards
func (list *LinkedList) ReversibleIterator(flipped bool) datastructures.Iterator {

	list.lock.RLock()
	defer list.lock.RUnlock()

	numberOfUpdatesAtCreation := list.numberOfUpdates
	return &concurrencyAwareIterator{list.innerList.ReversibleIterator(flipped), func() bool {
		return list.numberOfUpdates != numberOfUpdatesAtCreation
	}}
}

// Iterator is used to iterate over list values
func (list *LinkedList) Iterator() datastructures.Iterator {
	return list.ReversibleIterator(false)
}

// Add adds a new node to the tail of the list
func (list *LinkedList) Add(value interface{}) {

	list.lock.Lock()
	defer list.lock.Unlock()

	list.numberOfUpdates++

	list.innerList.Add(value)
}

// Clear clears the list to become empty
func (list *LinkedList) Clear() {

	list.lock.Lock()
	defer list.lock.Unlock()

	list.numberOfUpdates++

	list.innerList.Clear()
}

// Get value stored at specified index
func (list *LinkedList) Get(offset uint64, fromBack bool) interface{} {

	list.lock.RLock()
	defer list.lock.RUnlock()

	if list.innerList.Size() == 0 {
		return nil
	}

	return list.innerList.Get(offset, fromBack)
}

// Delete value stored at specified index
func (list *LinkedList) Delete(offset uint64, fromBack bool) interface{} {

	list.lock.Lock()
	defer list.lock.Unlock()

	if list.innerList.Size() == 0 {
		return nil
	}

	list.numberOfUpdates++

	return list.innerList.Delete(offset, fromBack)
}
