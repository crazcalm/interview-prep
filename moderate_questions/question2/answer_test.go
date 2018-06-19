package question2

import (
	"strings"
	"testing"
)

func TestTrimWord(t *testing.T) {
	tests := []struct {
		Input  string
		Expect []string
	}{
		{"", []string{}},
		{"I love life", []string{"i", "love", "life"}},
		{"I love life!", []string{"i", "love", "life"}},
		{"I love life?", []string{"i", "love", "life"}},
		{"I love life.", []string{"i", "love", "life"}},
		{"I    love    life!   ", []string{"i", "love", "life"}},
		{"I love life!    ", []string{"i", "love", "life"}},
	}

	for _, test := range tests {
		result := TrimWord(test.Input)

		if len(result) != len(test.Expect) {
			t.Errorf("The result has a different number of items than what was expected.\nExpected: %d (%v)\nResult: %d (%v)", len(test.Expect), test.Expect, len(result), result)
		}

		for index, value := range test.Expect {
			if !strings.EqualFold(value, result[index]) {
				t.Errorf("Expected %s, but got %s", value, result[index])
			}
		}
	}
}

func TestFrequency(t *testing.T) {
	tests := []struct {
		Input  []string
		Expect map[string]int
	}{
		{
			[]string{"i", "love", "life"},
			map[string]int{"i": 1, "love": 1, "life": 1},
		},
		{
			[]string{"a", "a", "a", "b", "d"},
			map[string]int{"a": 3, "b": 1, "d": 1},
		},
	}

	for _, test := range tests {
		result := Frequency(test.Input)
		if len(result) != len(test.Expect) {
			t.Errorf("The results has a different number of items that what was expected. Expected: %d (%v)\nResult: %d (%v)\n", len(test.Expect), test.Expect, len(result), result)
		}

		for key, value := range test.Expect {
			if value != result[key] {
				t.Errorf("Expected count of %s to be %d, but got %d", key, value, result[key])
			}
		}
	}
}

func TestWordFrequency(t *testing.T) {
	tests := []struct {
		Freq   map[string]int
		Word   string
		Count  int
		Total  int
		Answer float32
	}{
		{map[string]int{"a": 5, "b": 3, "c": 4}, "a", 5, 12, 0.41666666},
	}

	for _, test := range tests {
		count, total, result := WordFrequency(test.Freq, test.Word)

		if count != test.Count || total != test.Total || test.Answer != result {
			t.Errorf("Expected (count = %d) (total words = %d) and (result = %v), but got %d, %d, %v", test.Count, test.Total, test.Answer, count, total, result)
		}
	}
}

func TestFreqMerge(t *testing.T) {
	tests := []struct {
		Freq1  map[string]int
		Freq2  map[string]int
		Expect map[string]int
	}{
		{map[string]int{"a": 5, "b": 2}, map[string]int{"a": 7, "c": 3}, map[string]int{"a": 12, "b": 2, "c": 3}},
	}

	for _, test := range tests {
		result := FreqMerge(test.Freq1, test.Freq2)

		if len(test.Expect) != len(result) {
			t.Errorf("Not equal in length. Expected %d (%v), but got %d (%v)", len(test.Expect), test.Expect, len(result), result)
		}

		for key, value := range test.Expect {
			if result[key] != value {
				t.Errorf("Expected key %s to have count %d, but got %d", key, value, result[key])
			}
		}
	}
}
