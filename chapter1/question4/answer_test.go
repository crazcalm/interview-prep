package question4

import (
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		Input  string
		Answer bool
	}{
		{"tactcoa", true},
		{"abcdeft", false},
	}

	for _, test := range tests {
		result := answer(test.Input)
		if result != test.Answer {
			t.Errorf("Expected %t, but got %t", test.Answer, result)
		}
	}
}
