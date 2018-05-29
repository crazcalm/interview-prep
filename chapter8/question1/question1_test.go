package question1

import (
	"testing"
)

func testSteps(t *testing.T, function func(int) int) {
	tests := []struct {
		Input  int
		Answer int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 4},
		{4, 7},
		{5, 13},
	}

	for _, test := range tests {
		result := function(test.Input)
		if result != test.Answer {
			t.Errorf("Expected StepsRecursive(%d) = %d, but got %d", test.Input, test.Answer, result)
		}
	}
}

func TestStepsRecursion(t *testing.T) {
	testSteps(t, StepsRecursion)
}

func TestStepsMemoization(t *testing.T) {
	testSteps(t, StepsMemoization)
}

func benchmark32(b *testing.B, function func(int) int) {
	input := 32
	function(input)
}

func BenchmarkStepsRecursion(b *testing.B) {
	benchmark32(b, StepsRecursion)
}

func BenchmarkStepsMemoization(b *testing.B) {
	benchmark32(b, StepsMemoization)
}
