package queue

import (
	"container/list"
	"errors"
)

//Queue -- My queue interface
type Queue interface {
	Enqueue(data interface{})
	Dequeue() (data interface{}, err error)
	Len() int
}

//LinkedListQueue -- doubly linked list implementation
type LinkedListQueue struct {
	list *list.List
}

//NewLinkedListQueue -- returns an empty queue
func NewLinkedListQueue() Queue {
	list := list.New()
	return LinkedListQueue{list: list}
}

//Enqueue -- adds a new element to the Queue
func (q LinkedListQueue) Enqueue(data interface{}) {
	q.list.PushBack(data)
}

//Dequeue -- Removes the next element in the Queue
func (q LinkedListQueue) Dequeue() (data interface{}, err error) {
	if q.list.Len() == 0 {
		return data, errors.New("The queue is empty")
	}

	data = q.list.Remove(q.list.Front())

	return data, nil
}

//Len -- Returns the length of the Queue
func (q LinkedListQueue) Len() int {
	return q.list.Len()
}
