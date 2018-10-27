package concurrent

import (
	"github.com/saharshsingh/gophinoff/datastructures"
)

type concurrencyAwareIterator struct {
	wrappedIterator     datastructures.Iterator
	hasListBeenModified func() bool
}

func (iterator *concurrencyAwareIterator) HasNext() bool {

	if iterator.hasListBeenModified() {
		panic("List being iterated has been modified since iterator creation")
	}

	return iterator.wrappedIterator.HasNext()
}

func (iterator *concurrencyAwareIterator) Next() interface{} {

	if iterator.hasListBeenModified() {
		panic("List being iterated has been modified since iterator creation")
	}

	return iterator.wrappedIterator.Next()
}
