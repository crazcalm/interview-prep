package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		List   []int
		Value  int
		Expect int
	}{
		{[]int{}, 10, -1},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, 2},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8, 7},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7}, 4, 4},
	}

	for _, test := range tests {
		result := BinarySearch(test.List, test.Value)
		if result != test.Expect {
			t.Errorf("For list %v and value %d, expected answer to be %d, but got %d", test.List, test.Value, test.Expect, result)
		}
	}
}
