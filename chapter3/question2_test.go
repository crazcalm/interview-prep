package question2

import (
	"testing"
)

func TestTwoStackCase1(t *testing.T) {
	stack := NewTwoStack()
	expected := 1

	for _, data := range []int{1, 2, 3, 4, 5, 6} {
		stack = stack.Push(data)
	}

	answer, err := stack.Min()
	if err != nil {
		t.Errorf("Error occurred while trying to get Min value: %s", err.Error())
	}

	if answer.(int) != expected {
		t.Errorf("Expected %d, but got %d", expected, answer.(int))
	}
}

func TestTwoStackCase2(t *testing.T) {
	stack := NewTwoStack()
	expected := -4

	for _, data := range []int{1, 2, 3, 4, -4, 6} {
		stack = stack.Push(data)
	}

	answer, err := stack.Min()
	if err != nil {
		t.Errorf("Error occurred while trying to get Min value: %s", err.Error())
	}

	if answer.(int) != expected {
		t.Errorf("Expected %d, but got %d", expected, answer.(int))
	}
}

func TestTwoStackCase3(t *testing.T) {
	stack := NewTwoStack()
	expected := -6

	for _, data := range []int{1, -2, -3, -4, -5, -6} {
		stack = stack.Push(data)
	}

	answer, err := stack.Min()
	if err != nil {
		t.Errorf("Error occurred while trying to get Min value: %s", err.Error())
	}

	if answer.(int) != expected {
		t.Errorf("Expected %d, but got %d", expected, answer.(int))
	}
}

func TestTwoStackCase4(t *testing.T) {
	stack := NewTwoStack()
	expected := -5

	for _, data := range []int{1, -2, -3, 4, -5, -6} {
		stack = stack.Push(data)
	}

	stack, _, err := stack.Pop()
	if err != nil {
		t.Errorf("Unexpected error when popping the stack: %s", err.Error())
	}

	answer, err := stack.Min()
	if err != nil {
		t.Errorf("Error occurred while trying to get Min value: %s", err.Error())
	}

	if answer.(int) != expected {
		t.Errorf("Expected %d, but got %d", expected, answer.(int))
	}
}
