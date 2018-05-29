package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		Input  []int
		Expect []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{2, 1, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 1, 9, 9, 2, 2, 8, 8, 3, 3, 7, 7, 4, 4, 6, 6, 5, 5}, []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}},
	}

	for _, test := range tests {
		result := BubbleSort(test.Input)

		if len(result) != len(test.Expect) {
			t.Errorf("Expected %v, but got %v", test.Expect, result)
		}

		for i := 0; i < len(test.Expect); i++ {
			if result[i] != test.Expect[i] {
				t.Errorf("Expected %v, but got %v", test.Expect, result)
			}

		}
	}
}
