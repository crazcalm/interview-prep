package sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		Array    []int
		Solution []int
	}{
		{Array: []int{1, 8, 2, 5, 3, 6, 7, 9, 4}, Solution: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, test := range tests {
		QuickSort(test.Array)

		for i := 0; i < len(test.Array); i++ {
			if test.Array[i] != test.Solution[i] {
				t.Errorf("Expected sorted answer to be %v, but got %v", test.Solution, test.Array)

			}
		}
	}
}

func TestQuickSortString(t *testing.T) {
	tests := []struct {
		Array    []string
		Solution []string
	}{
		{Array: []string{"a", "g", "b", "f", "d", "e", "c"}, Solution: []string{"a", "b", "c", "d", "e", "f", "g"}},
	}

	for _, test := range tests {
		QuickSortString(test.Array)

		for i := 0; i < len(test.Array); i++ {
			if test.Array[i] != test.Solution[i] {
				t.Errorf("Expected sorted answer to be %v, but got %v", test.Solution, test.Array)

			}
		}
	}
}

func benchmark(b *testing.B, list []int, function func([]int)) {
	function(list)
}

func Benchmark10000000(b *testing.B) {
	numOfItems := 10000000
	list := make([]int, numOfItems)
	for i := 0; i < numOfItems; i++ {
		list[i] = i
	}

	benchmark(b, list, QuickSort)
}

func Benchmark1000000(b *testing.B) {
	numOfItems := 1000000
	list := make([]int, numOfItems)
	for i := 0; i < numOfItems; i++ {
		list[i] = i
	}

	benchmark(b, list, QuickSort)
}

func Benchmark100000(b *testing.B) {
	numOfItems := 100000
	list := make([]int, numOfItems)
	for i := 0; i < numOfItems; i++ {
		list[i] = i
	}

	benchmark(b, list, QuickSort)
}

func Benchmark10000(b *testing.B) {
	numOfItems := 10000
	list := make([]int, numOfItems)
	for i := 0; i < numOfItems; i++ {
		list[i] = i
	}

	benchmark(b, list, QuickSort)
}

func Benchmark1000(b *testing.B) {
	numOfItems := 1000
	list := make([]int, numOfItems)
	for i := 0; i < numOfItems; i++ {
		list[i] = i
	}

	benchmark(b, list, QuickSort)
}
