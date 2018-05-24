package trie

import (
	"strings"
)

//Node -- Node of a trie
type Node struct {
	Key      rune
	Value    interface{}
	Terminal bool
	Parent   *Node
	Children map[rune]*Node
}

//Trie -- The trie data structure
type Trie struct {
	Root *Node
}

//New -- Creates a new Trie
func New() Trie {
	node := new(Node)
	node.Children = make(map[rune]*Node)
	return Trie{Root: node}
}

func reverseSlice(list []string) []string {
	for i := 0; i < (len(list))/2; i++ {
		list[i], list[len(list)-1-i] = list[len(list)-1-i], list[i]
	}
	return list
}

//String -- returns the string associate with that node
func String(node *Node) string {
	var result []string
	for node.Parent != nil {
		result = append(result, string(node.Key))
		node = node.Parent
	}
	result = reverseSlice(result)
	return strings.Join(result, "")
}

//Depth -- Returns the depth of a node from the root
func Depth(node *Node) int {
	if node.Parent == nil {
		return 0
	}
	return Depth(node.Parent) + 1
}

//Find -- Returns the node associated with the end of the word.
func (t Trie) Find(word string) *Node {
	node := t.Root
	for _, char := range word {
		tempt, ok := node.Children[char]
		if !ok {
			return nil
		}
		node = tempt
	}
	return node
}

//Push -- Addes a node to the Trie
func (t Trie) Push(word string) Trie {
	var addNewNodes bool
	node := t.Root
	var lastIndexRune int
	for index, char := range word {
		tempt, ok := node.Children[char]
		if !ok {
			lastIndexRune = index
			addNewNodes = true
			break
		}

		//Move to the next node in the Trie
		node = tempt
	}

	//Add the non-existent chars to the trie
	if addNewNodes {
		for _, char := range word[lastIndexRune:] {
			tempt := new(Node)
			tempt.Key = char
			tempt.Parent = node
			tempt.Children = make(map[rune]*Node)
			node.Children[char] = tempt
			node = node.Children[char]
		}
	}
	node.Terminal = true
	return t
}

//Delete -- removes a word from the trie
func (t Trie) Delete(word string) {
	node := t.Find(word)
	if node == nil { //Word does not exist in trie
		return
	}

	//If node was a terminal node, then that needs to be changed
	if node.Terminal {
		node.Terminal = false
	}

	for i := len(word) - 1; i != 0; i-- {

		//If node has zero children, then the node needs to be removed
		if len(node.Children) == 0 {
			delete(node.Parent.Children, node.Key)
		}
		if node.Parent == nil { //reached root
			break
		}
		node = node.Parent

	}
}
