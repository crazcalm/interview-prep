package stack

import (
	"fmt"
	"strings"
	"testing"
)

func testNewStack(t *testing.T, function func() Stack) {
	stack := function()

	result := fmt.Sprintf("%T", stack)

	if !strings.Contains(result, "stack") {
		t.Errorf("Expected type Stack, but got %s", result)
	}
}

func testEmpty(t *testing.T, function func() Stack) {
	tests := []struct {
		Data   []int
		Expect bool
	}{
		{[]int{}, true},
		{[]int{1, 2, 3}, false},
	}

	for _, test := range tests {
		stack := function()
		for _, data := range test.Data {
			stack = stack.Push(data)
		}

		if stack.Empty() != test.Expect {
			t.Errorf("Expected stack.Empty() to return %t, but got %t", test.Expect, stack.Empty())
		}
	}
}

func testPeek(t *testing.T, function func() Stack) {
	tests := []struct {
		Data []int
		Err  bool
	}{
		{[]int{}, true},
		{[]int{1, 2, 3, 4, 5}, false},
	}

	for _, test := range tests {
		stack := function()
		for _, data := range test.Data {
			stack = stack.Push(data)
		}

		//Case: stack is empty
		if stack.Length() == 0 {
			_, err := stack.Peek()
			if test.Err && err != nil {
				continue
			} else {
				t.Errorf("Expected an error, but no error was received.")
			}
		}

		//Case: stack is not empty
		data, _ := stack.Peek()
		expect := test.Data[len(test.Data)-1]
		if data.(int) != expect {
			t.Errorf("Expected Peek data to be %d, but got %d", expect, data.(int))
		}
	}
}

func testPop(t *testing.T, function func() Stack) {
	tests := []struct {
		Data []int
		Err  bool
	}{
		{[]int{}, true},
		{[]int{1, 2, 3, 4, 5}, false},
	}

	for _, test := range tests {
		stack := NewLinkedListStack()
		for _, data := range test.Data {
			stack = stack.Push(data)
		}

		if stack.Length() == 0 {
			var err error
			stack, _, err = stack.Pop()
			if test.Err && err != nil {
				continue
			} else {
				t.Errorf("Expected an error, but no error was received.")
			}

		}

		for i := len(test.Data) - 1; i > 0; i-- {
			var data interface{}
			var err error
			stack, data, err = stack.Pop()
			if err != nil {
				t.Errorf("Unexpected err: %s", err.Error())
			}

			if data.(int) != test.Data[i] {
				t.Errorf("Expected %d (index %d) to be popped, but got %d", test.Data[i], i, data.(int))
			}
		}
	}
}

func testPush(t *testing.T, function func() Stack) {
	tests := []struct {
		Data   []int
		Length int
	}{
		{[]int{}, 0},
		{[]int{1, 2, 3}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9},
	}

	for _, test := range tests {
		stack := function()
		for _, data := range test.Data {
			stack = stack.Push(data)
		}

		if stack.Length() != test.Length {
			t.Errorf("Expected the length of the stack to be %d, but got %d", test.Length, stack.Length())
		}
	}
}

func benchmark(b *testing.B, function func() Stack) {
	stack := function()
	for i := 0; i < 1000000; i++ {
		stack.Push(i)
	}

	for !stack.Empty() {
		stack.Pop()
	}
}

func TestNewLinkedListStack(t *testing.T) {
	testNewStack(t, NewLinkedListStack)
}

func TestNewSliceStack(t *testing.T) {
	testNewStack(t, NewSliceStack)
}

func TestLinkedListStackEmpty(t *testing.T) {
	testEmpty(t, NewLinkedListStack)
}

func TestSliceStackEmpty(t *testing.T) {
	testEmpty(t, NewSliceStack)
}

func TestLinkedListStackPeek(t *testing.T) {
	testPeek(t, NewLinkedListStack)
}

func TestSliceStackPeek(t *testing.T) {
	testPeek(t, NewSliceStack)
}

func TestLinkedListStackPop(t *testing.T) {
	testPop(t, NewLinkedListStack)
}

func TestSliceStackPop(t *testing.T) {
	testPop(t, NewSliceStack)
}

func TestLinkedListStackPush(t *testing.T) {
	testPush(t, NewLinkedListStack)
}

func TestSliceStackPush(t *testing.T) {
	testPush(t, NewSliceStack)
}

func BenchmarkLinkedListStack(b *testing.B) {
	benchmark(b, NewLinkedListStack)
}

func BenchmarkSliceStack(b *testing.B) {
	benchmark(b, NewSliceStack)
}
