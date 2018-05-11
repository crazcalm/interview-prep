package question2

import (
	"testing"
)

func TestHash(t *testing.T) {
	tests := []struct {
		String1 string
		String2 string
		Expect  bool
	}{
		{"123456789", "123456789", true},
		{"123456789", "987654321", true},
		{"123", "1234", false},
		{"123", "124", false},
	}

	for _, test := range tests {
		answer := Hash(test.String1, test.String2)

		if answer != test.Expect {
			t.Errorf("For strings %s and %s, Expect %t, but got %t", test.String1, test.String2, test.Expect, answer)
		}
	}
}
