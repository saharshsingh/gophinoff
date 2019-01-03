package queue

import (
	"github.com/saharshsingh/gophinoff/datastructures"
	"github.com/saharshsingh/gophinoff/datastructures/linkedlist"
)

// Queue implements standard FIFO queue
type Queue struct {
	doublyLinkedList datastructures.LinkedList
}

// Create an empty queue
func Create() Queue {
	return Queue{linkedlist.Create()}
}

// Queue adds new element to tail of queue
func (queue *Queue) Queue(value interface{}) {
	queue.doublyLinkedList.Add(value)
}

// Peek the top of queue
func (queue *Queue) Peek() interface{} {
	if queue.Size() < 1 {
		return nil
	}
	return queue.doublyLinkedList.Get(0, false)
}

// Dequeue removes and returns the top of queue
func (queue *Queue) Dequeue() interface{} {
	if queue.Size() < 1 {
		return nil
	}
	return queue.doublyLinkedList.Delete(0, false)
}

// Size returns number of elements currently in queue
func (queue *Queue) Size() uint64 {
	return queue.doublyLinkedList.Size()
}
