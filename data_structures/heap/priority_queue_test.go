package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	items := map[string]int{
		"Your Family":       10,
		"Work":              7,
		"Significant Other": 9,
		"Children":          9,
		"friends":           5,
		"Co-workers":        3,
		"Strangers":         1,
	}

	index := 0
	pq := make(PriorityQueue, len(items))
	for value, priority := range items {
		pq[index] = &Item{
			index:    index,
			Value:    value,
			priority: priority,
		}
		index++
	}
	heap.Init(&pq)

	item := &Item{
		Value:    "Personal Care",
		priority: 8,
	}

	heap.Push(&pq, item)
	pq.Update(item, item.Value, 9)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%d: %v\n", item.priority, item.Value)
	}
}
