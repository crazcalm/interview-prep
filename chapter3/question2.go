package question2

import (
	"github.com/crazcalm/interview-prep/data_structures/stack"
)

/*
Stack Min: How would you design a stack which, in addition to push and pop, has a function min
		   which returns the minimum element? Push, pop, and min should all operate in O(1) time.
*/

/*
Note: My solution assumes that all the data is of type int
*/

//TwoStack -- A single stack that is built using two stacks
type TwoStack struct {
	stack1 stack.Stack
	stack2 stack.Stack
	min    interface{}
}

/*
type Stack interface {
	Pop() (Stack, interface{}, error)
	Push(data interface{}) Stack
	Peek() (data interface{}, err error)
	Empty() bool
	Length() int
}
*/

//NewTwoStack -- returns a new TwoStack
func NewTwoStack() TwoStack {
	list1 := stack.NewLinkedListStack()
	list2 := stack.NewLinkedListStack()
	min := 0
	return TwoStack{stack1: list1, stack2: list2, min: min}
}

//Push -- All elements get pushed to stack1. If the element is less than
//all other elements in the stack, then that element is also pushed to stack2
func (s TwoStack) Push(data interface{}) TwoStack {
	//Case: Stack is empty. Need to seed the current min
	if s.stack1.Length() == 0 {
		s.stack1 = s.stack1.Push(data)
		s.stack2 = s.stack2.Push(data)
		s.min = data
		return s
	}

	//Case: Stack is not empty

	//Add data to stack1
	s.stack1 = s.stack1.Push(data)

	//Check if data is a new min
	if data.(int) <= s.min.(int) {
		s.stack2 = s.stack2.Push(data)
		s.min = data
	}
	return s
}

//Pop -- Pops the last element off the stack
func (s TwoStack) Pop() (TwoStack, interface{}, error) {
	var data interface{}
	var err error

	s.stack1, data, err = s.stack1.Pop()
	if err != nil { //Case: Stack is empty
		return s, data, err
	}

	//Check to see if the data popped is equal to the current min
	if data.(int) == s.min.(int) {
		s.stack2, data, err = s.stack2.Pop()
		if err != nil {
			return s, data, err
		}

		//set new min
		data2, err := s.stack2.Peek()
		if err != nil {
			return s, data2, err
		}

		s.min = data2
	}

	return s, data, nil
}

//Peek -- Previews the top element
func (s TwoStack) Peek() (data interface{}, err error) {
	return s.stack1.Peek()
}

//Empty -- If empty, return true. Otherwise, return false
func (s TwoStack) Empty() bool {
	return s.stack1.Empty()
}

//Length -- Length of stack1
func (s TwoStack) Length() int {
	return s.stack1.Length()
}

//Min -- returns the data of the smallest element in the stack
func (s TwoStack) Min() (interface{}, error) {
	return s.stack2.Peek()
}
