package linkedlist

type linkedListIterator struct {
	current *node
	head    *node
	flipped bool
}

func (iter *linkedListIterator) HasNext() bool {
	return iter.current != nil
}

func (iter *linkedListIterator) Next() interface{} {
	return iter.next().value
}

func (iter *linkedListIterator) next() *node {

	current := iter.current

	nextIsHead := current == nil || current.next == iter.head
	if iter.flipped {
		nextIsHead = current == nil || current.prev == iter.head
	}

	if nextIsHead {
		iter.current = nil
	} else {
		if iter.flipped {
			iter.current = current.prev
		} else {
			iter.current = current.next
		}
	}

	return current
}
