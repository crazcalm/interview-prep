package question1

import (
	"testing"
)

const (
	BenchMarkCase = "abcdefghijklmnopqrstuvwxyz z"
)

var (
	Tests = []struct {
		String string
		Expect bool
	}{
		{"", true},
		{"a", true},
		{"abcdefghijklmnop", true},
		{"aa", false},
		{"abcdefghijklmnopf", false},
		{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ Z", false},
	}
)

func helper(t *testing.T, function func(string) bool) {
	for _, test := range Tests {
		answer := function(test.String)

		if answer != test.Expect {
			t.Errorf("For '%s', expected %t, but got %t", test.String, test.Expect, answer)
		}
	}
}

func TestStandardLibrary(t *testing.T) {
	helper(t, StandardLibrary)
}

func TestQuickSort(t *testing.T) {
	helper(t, QuickSort)
}

func TestHash(t *testing.T) {
	helper(t, Hash)
}

func TestBruteForce(t *testing.T) {
	helper(t, BruteForce)
}

func BenchmarkStandardLibrary(b *testing.B) {
	StandardLibrary(BenchMarkCase)
}

func BenchmarkQuickSort(b *testing.B) {
	QuickSort(BenchMarkCase)
}

func BenchmarkHash(b *testing.B) {
	Hash(BenchMarkCase)
}

func BenchmarkBruteForce(b *testing.B) {
	BruteForce(BenchMarkCase)
}
