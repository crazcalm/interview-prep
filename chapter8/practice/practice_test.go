package practice

import (
	"testing"
)

/*
Fib: Fib(n) = Fib(n - 1) + Fib(n - 2) where Fib(0) = 0 and Fib(1) = 1

Fib(0) = 0
Fib(1) = 1
Fib(2) = 1
Fib(3) = 2
Fib(4) = 3
Fib(5) = 5
Fib(6) = 8
Fib(7) = 13
Fib(8) = 21
*/

func testFib(t *testing.T, function func(int) int) {
	tests := []struct {
		Input  int
		Answer int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
	}

	for _, test := range tests {
		result := function(test.Input)
		if result != test.Answer {
			t.Errorf("Expected Fib(%d) = %d, but got %d", test.Input, test.Answer, result)
		}
	}
}

func TestRecursiveFib(t *testing.T) {
	testFib(t, RecursiveFib)
}

func TestInterativeFib(t *testing.T) {
	testFib(t, IterativeFib)
}

func TestTopDownMemoizationFib(t *testing.T) {
	testFib(t, TopDownMemoizationFib)
}

func TestBottomUpMemoizationFib(t *testing.T) {
	testFib(t, BottomUpMemoizationFib)
}

func benchmarkFib50(b *testing.B, function func(int) int) {
	input := 50
	function(input)
}

func benchmarkFib1000000(b *testing.B, function func(int) int) {
	input := 1000000
	function(input)
}

func BenchmarkRecursiveFib50(b *testing.B) {
	benchmarkFib50(b, RecursiveFib)
}

func BenchmarkIterativeFib50(b *testing.B) {
	benchmarkFib50(b, IterativeFib)
}

func BenchmarkTopDownMemoizationFib50(b *testing.B) {
	benchmarkFib50(b, TopDownMemoizationFib)
}

func BenchmarkIterativeFib10000000(b *testing.B) {
	benchmarkFib1000000(b, IterativeFib)
}

func BenchmarkTopDownMemoizationFib1000000(b *testing.B) {
	benchmarkFib1000000(b, TopDownMemoizationFib)
}

func BenchmarkBottomUpMemoizationFib1000000(b *testing.B) {
	benchmarkFib1000000(b, BottomUpMemoizationFib)
}
