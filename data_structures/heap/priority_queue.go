package heap

import (
	"container/heap"
)

//Item -- the items held in the priority queue
type Item struct {
	index    int         // Reference to the index within the queue
	Value    interface{} // Payload
	priority int         //  Priority weight
}

//PriorityQueue -- Generic PriorityQueue
type PriorityQueue []*Item

//Len -- returns the length of the Priority Queue
func (pq PriorityQueue) Len() int { return len(pq) }

//Less -- Defines what it means for one item to be less than another
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority > pq[j].priority }

//Swap -- Defines what it means to swap two items in the Priority Queue
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

//Push -- Pushes an item to the priority queue
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

//Pop -- Pops the item with the highest priority from the queue
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // For safety
	*pq = old[0 : n-1]
	return item
}

//Update -- Update an item in the queue
func (pq *PriorityQueue) Update(item *Item, value interface{}, priority int) {
	item.Value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
