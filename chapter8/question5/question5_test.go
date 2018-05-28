package question5

import (
	"testing"
)

func testSolution(t *testing.T, function func(uint32, uint32) uint32) {
	tests := []struct {
		A      uint32
		B      uint32
		Answer uint32
	}{
		{1, 9, 9},
		{2, 9, 18},
		{3, 9, 27},
		{4, 9, 36},
		{5, 9, 45},
		{6, 9, 54},
		{7, 9, 63},
		{8, 9, 72},
		{9, 9, 81},
		{9, 1, 9},
		{9, 2, 18},
		{9, 3, 27},
		{9, 4, 36},
		{9, 5, 45},
		{9, 6, 54},
		{9, 7, 63},
		{9, 8, 72},
	}

	for _, test := range tests {
		result := function(test.A, test.B)
		if result != test.Answer {
			t.Errorf("Expected %d * %d = %d, but got %d", test.A, test.B, test.Answer, result)
		}
	}
}

func TestSimpleRecursion(t *testing.T) {
	testSolution(t, SimpleRecursion)
}

func TestSlightlyBetter(t *testing.T) {
	testSolution(t, SlightlyBetter)
}

func TestSimpleIteration(t *testing.T) {
	testSolution(t, SimpleIteration)
}

func TestBetter(t *testing.T) {
	testSolution(t, Better)
}

func benchmarkSolution(b *testing.B, function func(uint32, uint32) uint32) {
	var a, c uint32
	a = 10000000
	c = 10000000
	function(a, c)
}

func benchmarkSolution2(b *testing.B, function func(uint32, uint32) uint32) {
	var a, c uint32
	a = 100000000
	c = 100000000
	function(a, c)
}

func BenchmarkSimpleRecursion10000000(b *testing.B) {
	benchmarkSolution(b, SimpleRecursion)
}

func BenchmarkSlightlyBetter10000000(b *testing.B) {
	benchmarkSolution(b, SlightlyBetter)
}

func BenchmarkBetter10000000(b *testing.B) {
	benchmarkSolution(b, Better)
}

func BenchmarkSimpleIteration10000000(b *testing.B) {
	benchmarkSolution(b, SimpleIteration)
}

func BenchmarkBetter100000000(b *testing.B) {
	benchmarkSolution2(b, Better)
}

func BenchmarkSimpleIteration100000000(b *testing.B) {
	benchmarkSolution2(b, SimpleIteration)
}
