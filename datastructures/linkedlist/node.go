package linkedlist

type node struct {
	value interface{}
	next  *node
	prev  *node
}

func (n *node) prepend(value interface{}) *node {
	toPrepend := &node{value, n, n.prev}
	n.prev.next, n.prev = toPrepend, toPrepend
	return toPrepend
}

func (n *node) delete() interface{} {

	n.next.prev, n.prev.next = n.prev, n.next

	value := n.value
	n.value, n.next, n.prev = nil, nil, nil

	return value
}
