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
