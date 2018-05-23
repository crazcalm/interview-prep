package heap

import (
	"container/heap"
	"testing"
)

func TestMinIntHeap(t *testing.T) {
	minHeap := &MinIntHeap{5, 3, 4, 1, 2}

	heap.Init(minHeap)
	heap.Push(minHeap, 6)
	heap.Push(minHeap, 0)

	var count int
	for minHeap.Len() > 0 {
		result := heap.Pop(minHeap)
		if count != result.(int) {
			t.Errorf("Expected %d, but got %d", count, result.(int))
		}
		count++
	}
}

func TestMaxIntHeap(t *testing.T) {
	maxHeap := &MaxIntHeap{5, 3, 4, 1, 2}

	heap.Init(maxHeap)
	heap.Push(maxHeap, 6)
	heap.Push(maxHeap, 0)

	count := 6
	for maxHeap.Len() > 0 {
		result := heap.Pop(maxHeap)
		if count != result.(int) {
			t.Errorf("Expected %d, but got %d", count, result.(int))
		}
		count--
	}
}
