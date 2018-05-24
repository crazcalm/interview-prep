package trie

import (
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	_ = New()
}

func TestDelete(t *testing.T) {
	tests := []struct {
		Words      []string
		Delete     string
		NilList    []string
		NotNilList []string
	}{
		{[]string{"Happy", "Hello", "Help", "Band"}, "Hello", []string{"Hello", "Hell"}, []string{"Hel", "Help", "Happy", "Band"}},
		{[]string{"Happy", "Hello", "Help", "Band", "Helped"}, "Help", []string{}, []string{"Hel", "Help", "Happy", "Band", "Hello", "Helped"}},
	}

	for _, test := range tests {
		trie := New()
		for _, word := range test.Words {
			trie.Push(word)
		}

		trie.Delete(test.Delete)

		for _, word := range test.NilList {
			if node := trie.Find(word); node != nil {
				t.Errorf("Expected node to be nil for word %s", word)
			}
		}

		for _, word := range test.NotNilList {
			node := trie.Find(word)

			if node == nil {
				t.Errorf("Expected node to not be nil for word %s", word)
			}

			if strings.EqualFold(word, test.Delete) {
				if node.Terminal {
					t.Errorf("Word %s was deleted. This node should be set as not terminal: node -- %v", word, node)
				}
			}

		}
	}
}

func TestDepth(t *testing.T) {
	tests := []struct {
		Words  []string
		Find   string
		Expect int
	}{
		{[]string{"Happy", "Hello", "Help"}, "Happy", 5},
		{[]string{"Happy", "Hello", "Help"}, "Help", 4},
		{[]string{"Happy", "Hello", "Help"}, "H", 1},
	}

	for _, test := range tests {
		trie := New()

		for _, word := range test.Words {
			trie.Push(word)
		}

		node := trie.Find(test.Find)
		if node == nil {
			t.Errorf("Failed trying to find node associated with %s", test.Find)
		}
		result := Depth(node)
		if result != test.Expect {
			t.Errorf("Expected %d, but got %d", test.Expect, result)
		}
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		Words  []string
		Find   string
		Expect string
	}{
		{[]string{"Happy", "Hello", "Help"}, "Happy", "Happy"},
		{[]string{"Happy", "Hello", "Help"}, "Help", "Help"},
		{[]string{"Happy", "Hello", "Help"}, "Hel", "Hel"},
	}

	for _, test := range tests {
		trie := New()

		for _, word := range test.Words {
			trie.Push(word)
		}

		node := trie.Find(test.Find)
		if node == nil {
			t.Errorf("Failed trying to find node associated with %s", test.Find)
		}

		result := String(node)
		if !strings.EqualFold(result, test.Expect) {
			t.Errorf("Expected %s, but got %s", test.Expect, result)
		}
	}
}

func TestPush(t *testing.T) {
	trie := New()

	trie.Push("Happy")
	trie.Push("Hello")
	trie.Push("Help")
}

func TestFind(t *testing.T) {
	trie := New()

	trie.Push("Happy")
	trie.Push("Hello")
	trie.Push("Help")

	if result := trie.Find("Does not exist"); result != nil {
		t.Error("Expected this to return nil")
	}
}
