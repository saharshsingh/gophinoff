package datastructures

// Iterator interface for various types
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// LinkedList interface describing a circular doubly linked list
type LinkedList interface {
	Size() uint64
	ReversibleIterator(flipped bool) Iterator
	Iterator() Iterator
	Add(value interface{})
	Clear()
	Get(offset uint64, fromBack bool) interface{}
	Delete(offset uint64, fromBack bool) interface{}
}

// Comparable describes things that can be compared. Implementations of Compare() should return:
//
//	- 0 for equality
//	- negative integer if receiver is less than argument
//	- positive integer if receiver is greater than argument
type Comparable interface {
	Compare(other Comparable) int
}
