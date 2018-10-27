package stack

import (
	"github.com/saharshsingh/gophinoff/datastructures"
	"github.com/saharshsingh/gophinoff/datastructures/linkedlist"
)

// Stack implements standard LIFO stack
type Stack struct {
	doublyLinkedList datastructures.LinkedList
}

// Create an empty stack
func Create() Stack {
	return Stack{linkedlist.Create()}
}

// Push adds new element to head of stack
func (stack *Stack) Push(value interface{}) {
	stack.doublyLinkedList.Add(value)
}

// Peek the top of stack
func (stack *Stack) Peek() interface{} {
	if stack.Size() < 1 {
		return nil
	}
	return stack.doublyLinkedList.Get(0, true)
}

// Pop removes and returns the top of stack
func (stack *Stack) Pop() interface{} {
	if stack.Size() < 1 {
		return nil
	}
	return stack.doublyLinkedList.Delete(0, true)
}

// Size returns number of elements currently in stack
func (stack *Stack) Size() uint64 {
	return stack.doublyLinkedList.Size()
}
