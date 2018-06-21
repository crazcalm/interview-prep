package bst

import (
	"fmt"
	"testing"
)

func TestNewTree(t *testing.T) {
	_ = NewTree()
}

func TestRecursiveMin(t *testing.T) {
	bst := NewTree()

	data := []uint32{5, 6, 4, 7, 3, 8, 2, 10, 1}
	var expected uint32 = 1

	for _, v := range data {
		element := &Element{Key: v, Value: v, Parent: nil, Right: nil, Left: nil}
		bst.Insert(element)
	}

	result := RecursiveMin(bst.Root)
	if result.Value.(uint32) != expected {
		t.Errorf("Expected %d, but got %v", expected, result)
	}

}

func TestMin(t *testing.T) {
	bst := NewTree()

	data := []uint32{5, 6, 4, 7, 3, 8, 2, 10, 1}
	var expected uint32 = 1

	for _, v := range data {
		element := &Element{Key: v, Value: v, Parent: nil, Right: nil, Left: nil}
		bst.Insert(element)
	}

	result := bst.Min(bst.Root)
	if result.Value.(uint32) != expected {
		t.Errorf("Expected %d, but got %v", expected, result)
	}

}

func TestRecursiveInsert(t *testing.T) {
	bst := NewTree()

	data := []uint32{3, 5, 4, 2, 1}
	for _, v := range data {
		element := &Element{Key: v, Value: v, Parent: nil, Right: nil, Left: nil}
		bst = RecursiveInsert(bst.Root, element, bst)
	}

	for _, v := range data {
		result := RecursiveSearch(bst.Root, v)

		if v != result.Value.(uint32) {
			t.Errorf("Expected node to have Value = %d, but got %v", v, result)
		}
	}
}

func TestInsert(t *testing.T) {
	bst := NewTree()

	data := []uint32{1, 2, 3, 4, 5}
	for _, v := range data {
		element := &Element{Key: v, Value: v, Parent: nil, Right: nil, Left: nil}
		bst.Insert(element)
	}

	for _, v := range data {
		result := bst.Search(v)

		if v != result.Value.(uint32) {
			t.Errorf("Expected node to have Value = %d, but got %v", v, result)
		}
	}
}

func TestRecursiveSearch(t *testing.T) {
	bst := NewTree()

	data := []uint32{1, 2, 3, 4, 5}
	for _, v := range data {
		element := &Element{Key: v, Value: v, Parent: nil, Right: nil, Left: nil}
		bst.Insert(element)
	}

	for _, v := range data {
		result := RecursiveSearch(bst.Root, v)

		if v != result.Value.(uint32) {
			t.Errorf("Expected node to have Value = %d, but got %v", v, result)
		}
	}
}

func TestDelete(t *testing.T) {
	bst := NewTree()

	data := []uint32{1, 2, 3, 4, 5}
	for _, v := range data {
		element := &Element{Key: v, Value: v, Parent: nil, Right: nil, Left: nil}
		bst.Insert(element)
	}

	for _, v := range data {
		result := bst.Delete(v)
		if result == nil {
			t.Errorf("Failed to delete element with key %d", v)
		}
	}

	//Tree should be empty
	if bst.Root != nil {
		t.Errorf("Expected the tree to be empty, but it was not: %v", bst)
	}
}

func TestTraverseInOrder(t *testing.T) {
	bst := NewTree()

	data := []uint32{5, 4, 3, 2, 1}
	for _, v := range data {
		element := &Element{Key: v, Value: v, Parent: nil, Right: nil, Left: nil}
		bst.Insert(element)
	}
	callback := func(node *Element) {
		fmt.Printf("Node: %v\n", node)
	}

	TraverseInOrder(bst.Root, callback)
}

func TestTraverseReverseOrder(t *testing.T) {
	bst := NewTree()

	data := []uint32{5, 4, 3, 2, 1}
	for _, v := range data {
		element := &Element{Key: v, Value: v, Parent: nil, Right: nil, Left: nil}
		bst.Insert(element)
	}
	callback := func(node *Element) {
		fmt.Printf("Node: %v\n", node)
	}

	TraverseReverseOrder(bst.Root, callback)
}
