package concurrent

import (
	"sync"
	"testing"

	"github.com/saharshsingh/gophinoff/datastructures"
	"github.com/saharshsingh/gophinoff/datastructures/linkedlist"
	"github.com/saharshsingh/gophinoff/testutils"
)

func TestSatisfiesLinkedListInterface(t *testing.T) {
	var list datastructures.LinkedList = CreateConcurrentList()
	testutils.AssertEquals("Expected Size() to be 0", uint64(0), list.Size(), t)
}

func TestConcurrentAdd(t *testing.T) {

	list := CreateConcurrentList()

	batchSize := 100000
	insertConcurrently(list, batchSize)

	testutils.AssertEquals("Expected Size() at end to be '800000'", uint64(8*batchSize), list.Size(), t)
}

func TestConcurrentIterators(t *testing.T) {

	unsafeSortedList := linkedlist.CreateSorted()
	list := WrapWithConcurrentList(unsafeSortedList)

	batchSize := 1000
	insertConcurrently(list, batchSize)

	testutils.AssertEquals("Expected Size() at end to be '800000'", uint64(8*batchSize), list.Size(), t)

	iterator := list.Iterator()
	iterationCount := 0
	for ; iterator.HasNext(); iterationCount++ {
		testutils.AssertEquals("Expected list to be sorted from 0 to 8*batchSize", iterationCount, iterator.Next(), t)
	}
	testutils.AssertEquals("Expected loop to iterate 8*batchSize times", 8*batchSize, iterationCount, t)

	reverseIterator := list.ReversibleIterator(true)
	reverseIterationCount := 8 * batchSize
	for ; reverseIterator.HasNext(); reverseIterationCount-- {
		testutils.AssertEquals("Expected list to be sorted from 8*batchSize to 0", reverseIterationCount-1, reverseIterator.Next(), t)
	}
	testutils.AssertEquals("Expected loop to iterate 8*batchSize times", 0, reverseIterationCount, t)
}

func TestIteratorHasNextForConcurrentModification(t *testing.T) {

	list := CreateConcurrentList()
	list.Add(1)

	iterator := list.Iterator()
	testutils.AssertTrue("Expected HasNext() to return true", iterator.HasNext(), t)

	list.Add(2)

	defer testutils.AssertPanic("Expected next statement to cause panic", t)
	iterator.HasNext()

}

func TestIteratorNextForConcurrentModification(t *testing.T) {

	list := CreateConcurrentList()
	list.Add(1)
	list.Add(2)

	iterator := list.Iterator()
	testutils.AssertEquals("Expected Next() to return 1", 1, iterator.Next(), t)

	list.Add(3)

	defer testutils.AssertPanic("Expected next statement to cause panic", t)
	iterator.Next()

}

func TestConcurrentGetAndDelete(t *testing.T) {

	list := CreateConcurrentList()

	for i := 0; i < 20; i++ {
		list.Add(i)
	}

	testutils.AssertEquals("Expected Size() at beginning to be '20'", uint64(20), list.Size(), t)

	getBatch := func(offset uint64, wg *sync.WaitGroup) {

		defer wg.Done()

		for i := 0; i < 10000; i++ {
			for j := offset; j < offset+uint64(5); j++ {
				testutils.AssertEquals("Get test failed", int(j), list.Get(j, false), t)
			}
		}
	}

	addAndDeleteOne := func(wg *sync.WaitGroup) {

		defer wg.Done()

		for i := 0; i < 10000; i++ {
			list.Add("ToDelete")
			list.Delete(20, false)
		}
	}

	var wg sync.WaitGroup
	wg.Add(8)

	go getBatch(0, &wg)
	go getBatch(5, &wg)
	go getBatch(10, &wg)
	go getBatch(15, &wg)
	go addAndDeleteOne(&wg)
	go addAndDeleteOne(&wg)
	go addAndDeleteOne(&wg)
	go addAndDeleteOne(&wg)

	wg.Wait()

	testutils.AssertEquals("Expected Size() at end to be '20'", uint64(20), list.Size(), t)
	for i := 0; i < 20; i++ {
		testutils.AssertEquals("Get test failed", i, list.Get(uint64(i), false), t)
	}

}

func TestConcurrentClear(t *testing.T) {

	list := CreateConcurrentList()

	const NumTasks int = 8
	const iterations int = 10000

	var clearPhaseInit sync.WaitGroup
	var clearPhase sync.WaitGroup
	var postClearPhaseInit sync.WaitGroup
	var postClearPhase sync.WaitGroup
	var tasks sync.WaitGroup

	addClearAdd := func() {

		defer tasks.Done()

		for i := 0; i < iterations; i++ {

			clearPhaseInit.Wait()
			list.Add("Should get cleared")
			list.Clear()
			clearPhase.Done()

			postClearPhaseInit.Wait()
			list.Add("Should NOT get cleared")
			postClearPhase.Done()
		}
	}

	// initialize
	tasks.Add(NumTasks)
	clearPhaseInit.Add(1)
	for i := 0; i < NumTasks; i++ {
		go addClearAdd()
	}

	// coordinate each iteration
	for i := 0; i < iterations; i++ {

		clearPhase.Add(NumTasks)
		postClearPhaseInit.Add(1) // Tasks will wait on this before begining post clear phase
		clearPhaseInit.Done()
		clearPhase.Wait()

		postClearPhase.Add(NumTasks)
		clearPhaseInit.Add(1) // Tasks will wait on this before begining clear phase
		postClearPhaseInit.Done()
		postClearPhase.Wait()
	}

	tasks.Wait()

	testutils.AssertEquals("Expected Size() at end to be same as number of tasks", uint64(NumTasks), list.Size(), t)
	for i := uint64(0); i < uint64(NumTasks); i++ {
		testutils.AssertEquals("Expected 'Should NOT get cleared'", "Should NOT get cleared", list.Get(i, false), t)
	}
}

func TestGetWhenEmpty(t *testing.T) {
	list := CreateConcurrentList()
	testutils.AssertEquals("Expected nil when list is empty", nil, list.Get(0, false), t)
	testutils.AssertEquals("Expected nil when list is empty", nil, list.Get(0, true), t)
}

func TestDeleteWhenEmpty(t *testing.T) {
	list := CreateConcurrentList()
	testutils.AssertEquals("Expected nil when list is empty", nil, list.Delete(0, false), t)
	testutils.AssertEquals("Expected nil when list is empty", nil, list.Delete(0, true), t)
}

func insertConcurrently(list *LinkedList, batchSize int) {

	insertBatch := func(offset int, count int, wg *sync.WaitGroup) {

		defer wg.Done()

		for i := offset; i < offset+count; i++ {
			list.Add(i)
		}
	}

	var wg sync.WaitGroup
	wg.Add(8)

	go insertBatch(0, batchSize, &wg)
	go insertBatch(batchSize, batchSize, &wg)
	go insertBatch(2*batchSize, batchSize, &wg)
	go insertBatch(3*batchSize, batchSize, &wg)
	go insertBatch(4*batchSize, batchSize, &wg)
	go insertBatch(5*batchSize, batchSize, &wg)
	go insertBatch(6*batchSize, batchSize, &wg)
	go insertBatch(7*batchSize, batchSize, &wg)

	wg.Wait()

}
