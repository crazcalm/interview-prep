package stack

import (
	"errors"

	"github.com/crazcalm/interview-prep/data_structures/linked_list"
)

/*
I will implement two versions of a stack:

- Linkedlist
- Array
*/

//Stack -- My interface for a stack
type Stack interface {
	Pop() (Stack, interface{}, error)
	Push(data interface{}) Stack
	Peek() (data interface{}, err error)
	Empty() bool
	Length() int
}

//SliceStack -- Stack implementation that uses a slice
type SliceStack struct {
	slice []interface{}
}

//NewSliceStack -- Returns an empty stack
func NewSliceStack() Stack {
	return SliceStack{slice: []interface{}{}}
}

//Push -- Adds an element to the stack
func (s SliceStack) Push(data interface{}) Stack {
	s.slice = append(s.slice, data)
	return s
}

//Pop -- Pops the last element off the stack
func (s SliceStack) Pop() (Stack, interface{}, error) {
	if len(s.slice) == 0 {
		return s, nil, errors.New("The Stack is empty")
	}
	data := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]

	return s, data, nil
}

//Peek -- Allows you to see the top element on the stack
func (s SliceStack) Peek() (data interface{}, err error) {
	if len(s.slice) == 0 {
		return data, errors.New("The Stack is empty")
	}
	return s.slice[len(s.slice)-1], nil
}

//Empty -- If stack is empty, returns true. Otherwise, returns false
func (s SliceStack) Empty() bool {
	if len(s.slice) == 0 {
		return true
	}
	return false
}

//Length -- Returns the number of elements in the stack
func (s SliceStack) Length() int {
	return len(s.slice)
}

//LinkedListStack -- Stack implementation that uses linkedlists
type LinkedListStack struct {
	list linkedlist.List
	size int
}

//NewLinkedListStack -- Returns a stack that uses a linked list implementation
func NewLinkedListStack() Stack {
	list := linkedlist.List{FirstNode: nil}
	return LinkedListStack{list: list, size: 0}
}

//Empty -- Returns true if the stack is empty. False, otherwise
func (s LinkedListStack) Empty() bool {
	if s.list.FirstNode == nil {
		return true
	}
	return false
}

//Pop -- Pops the last element to enter the stack
func (s LinkedListStack) Pop() (Stack, interface{}, error) {
	var data interface{}
	if s.Empty() {
		return s, data, errors.New("The Stack is Empty")
	}

	//Case 1: Popping the head of the linked list
	if s.list.FirstNode.Next == nil {
		data = s.list.FirstNode.Data
		s.list = linkedlist.RemoveBeginning(s.list)
	}

	//Case 2: Popping a Node that is not the head of the linked list
	for n := s.list.FirstNode; n != nil; n = n.Next {
		if n.Next.Next == nil {
			data = n.Next.Data
			n.Next = n.Next.Next
			break
		}
	}

	//Decriment size
	s.size--

	return s, data, nil
}

//Push -- Adds an element onto the stack
func (s LinkedListStack) Push(data interface{}) Stack {
	node := &linkedlist.Node{Data: data, Next: nil}
	if s.list.FirstNode == nil {
		s.list = linkedlist.InsertBeginning(s.list, node)
	} else {
		for n := s.list.FirstNode; n != nil; n = n.Next {
			if n.Next == nil {
				linkedlist.InsertAfter(n, node)
				break
			}
		}
	}

	//Incriment size
	s.size++

	return s
}

//Peek -- Allows you to see the top element on the stack
func (s LinkedListStack) Peek() (data interface{}, err error) {
	_, data, err = s.Pop()
	if err != nil {
		return
	}
	s.Push(data)
	return
}

//Length -- returns the number of elements in the stack
func (s LinkedListStack) Length() int {
	return s.size
}
