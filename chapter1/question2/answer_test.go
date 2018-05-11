package question2

import (
	"testing"
)

const (
	BenchMarkString1 = "11111111112222222222333333333344444444441qazxsw23edcvfr45tgbnhy67ujm,ki89ol."
	BenchMarkString2 = ".lo98ik,mju76yhnbgt54rfvcde32wsxzaq14444444444333333333322222222221111111111"
)

func helper(t *testing.T, function func(string, string) bool) {
	tests := []struct {
		String1 string
		String2 string
		Expect  bool
	}{
		{"123456789", "123456789", true},
		{"123456789", "987654321", true},
		{"123", "1234", false},
		{"123", "124", false},
		{BenchMarkString1, BenchMarkString2, true},
	}

	for _, test := range tests {
		answer := function(test.String1, test.String2)

		if answer != test.Expect {
			t.Errorf("For strings %s and %s, Expect %t, but got %t", test.String1, test.String2, test.Expect, answer)
		}
	}
}

func TestMergeSort(t *testing.T) {
	helper(t, MergeSort)
}

func TestHash(t *testing.T) {
	helper(t, Hash)
}

func BenchmarkHash(b *testing.B) {
	Hash(BenchMarkString1, BenchMarkString2)
}

func BenchmarkMergeSort(b *testing.B) {
	MergeSort(BenchMarkString1, BenchMarkString2)
}
