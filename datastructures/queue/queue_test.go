package queue

import (
	"testing"

	"github.com/saharshsingh/gophinoff/testutils"
)

type student struct {
	name  string
	grade int
}

func TestCreate(t *testing.T) {

	queue := Create()

	testutils.AssertEquals("Expected size to be '0'", uint64(0), queue.Size(), t)
	testutils.AssertEquals("Expected peek to be 'nil'", nil, queue.Peek(), t)
	testutils.AssertEquals("Expected peek to be 'nil'", nil, queue.Deque(), t)
}

func TestPeek(t *testing.T) {

	saharsh, shan, naina := &student{"Saharsh Singh", 95}, &student{"Shan Singh", 75}, &student{"Naina Singh", 100}

	queue := Create()
	queue.Queue(saharsh)
	queue.Queue(shan)
	queue.Queue(naina)

	testutils.AssertEquals("Expected 'saharsh'", saharsh, queue.Peek(), t)
	testutils.AssertEquals("Expected size to be '0'", uint64(3), queue.Size(), t)

	testutils.AssertEquals("Expected 'saharsh'", saharsh, queue.Peek(), t)
	testutils.AssertEquals("Expected size to be '0'", uint64(3), queue.Size(), t)

	testutils.AssertEquals("Expected 'saharsh'", saharsh, queue.Peek(), t)
	testutils.AssertEquals("Expected size to be '0'", uint64(3), queue.Size(), t)
}

func TestDequeue(t *testing.T) {

	saharsh, shan, naina := &student{"Saharsh Singh", 95}, &student{"Shan Singh", 75}, &student{"Naina Singh", 100}

	queue := Create()
	queue.Queue(saharsh)
	queue.Queue(shan)
	queue.Queue(naina)

	testutils.AssertEquals("Expected 'saharsh' from Peek()", saharsh, queue.Peek(), t)
	testutils.AssertEquals("Expected size to be '3'", uint64(3), queue.Size(), t)

	testutils.AssertEquals("Expected 'saharsh' from Deque()", saharsh, queue.Deque(), t)
	testutils.AssertEquals("Expected 'shan' from Peek()", shan, queue.Peek(), t)
	testutils.AssertEquals("Expected size to be '2'", uint64(2), queue.Size(), t)

	testutils.AssertEquals("Expected 'shan' from Deque()", shan, queue.Deque(), t)
	testutils.AssertEquals("Expected 'naina' from Peek()", naina, queue.Peek(), t)
	testutils.AssertEquals("Expected size to be '1'", uint64(1), queue.Size(), t)

	testutils.AssertEquals("Expected 'naina' from Deque()", naina, queue.Deque(), t)
	testutils.AssertEquals("Expected 'nil' from Peek()", nil, queue.Peek(), t)
	testutils.AssertEquals("Expected size to be '0'", uint64(0), queue.Size(), t)

	testutils.AssertEquals("Expected 'nil' from Deque()", nil, queue.Deque(), t)
	testutils.AssertEquals("Expected size to be '0'", uint64(0), queue.Size(), t)
}
