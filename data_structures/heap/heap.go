package heap

//MinIntHeap -- Min heap containing integers
type MinIntHeap []int

func (h MinIntHeap) Len() int           { return len(h) }
func (h MinIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

//Pop -- removes an element from the heap
func (h *MinIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//Push -- adds an element to the heap
func (h *MinIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

//MaxIntHeap -- Min heap containing integers
type MaxIntHeap []int

func (h MaxIntHeap) Len() int           { return len(h) }
func (h MaxIntHeap) Less(i, j int) bool { return h[i] >= h[j] }
func (h MaxIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

//Pop -- removes an element from the heap
func (h *MaxIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//Push -- adds an element to the heap
func (h *MaxIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
