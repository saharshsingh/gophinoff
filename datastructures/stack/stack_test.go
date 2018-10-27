package stack

import (
	"testing"

	"github.com/saharshsingh/gophinoff/testutils"
)

type student struct {
	name  string
	grade int
}

func TestCreate(t *testing.T) {

	stack := Create()

	testutils.AssertEquals("Expected size to be '0'", uint64(0), stack.Size(), t)
	testutils.AssertEquals("Expected peek to be 'nil'", nil, stack.Peek(), t)
	testutils.AssertEquals("Expected peek to be 'nil'", nil, stack.Pop(), t)
}

func TestPeek(t *testing.T) {

	saharsh, shan, naina := &student{"Saharsh Singh", 95}, &student{"Shan Singh", 75}, &student{"Naina Singh", 100}

	stack := Create()
	stack.Push(saharsh)
	stack.Push(shan)
	stack.Push(naina)

	testutils.AssertEquals("Expected 'naina'", naina, stack.Peek(), t)
	testutils.AssertEquals("Expected size to be '0'", uint64(3), stack.Size(), t)

	testutils.AssertEquals("Expected 'naina'", naina, stack.Peek(), t)
	testutils.AssertEquals("Expected size to be '0'", uint64(3), stack.Size(), t)

	testutils.AssertEquals("Expected 'naina'", naina, stack.Peek(), t)
	testutils.AssertEquals("Expected size to be '0'", uint64(3), stack.Size(), t)
}

func TestPop(t *testing.T) {

	saharsh, shan, naina := &student{"Saharsh Singh", 95}, &student{"Shan Singh", 75}, &student{"Naina Singh", 100}

	stack := Create()
	stack.Push(saharsh)
	stack.Push(shan)
	stack.Push(naina)

	testutils.AssertEquals("Expected 'naina' from Peek()", naina, stack.Peek(), t)
	testutils.AssertEquals("Expected size to be '3'", uint64(3), stack.Size(), t)

	testutils.AssertEquals("Expected 'naina' from Deque()", naina, stack.Pop(), t)
	testutils.AssertEquals("Expected 'shan' from Peek()", shan, stack.Peek(), t)
	testutils.AssertEquals("Expected size to be '2'", uint64(2), stack.Size(), t)

	testutils.AssertEquals("Expected 'shan' from Deque()", shan, stack.Pop(), t)
	testutils.AssertEquals("Expected 'saharsh' from Peek()", saharsh, stack.Peek(), t)
	testutils.AssertEquals("Expected size to be '1'", uint64(1), stack.Size(), t)

	testutils.AssertEquals("Expected 'saharsh' from Deque()", saharsh, stack.Pop(), t)
	testutils.AssertEquals("Expected 'nil' from Peek()", nil, stack.Peek(), t)
	testutils.AssertEquals("Expected size to be '0'", uint64(0), stack.Size(), t)

	testutils.AssertEquals("Expected 'nil' from Pop()", nil, stack.Pop(), t)
	testutils.AssertEquals("Expected size to be '0'", uint64(0), stack.Size(), t)
}
