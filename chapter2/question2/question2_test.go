package question2

import (
	"testing"

	"github.com/crazcalm/interview-prep/data_structures/linked_list"
)

func linkedListCase1() linkedlist.List {
	node1 := &linkedlist.Node{Data: 1, Next: nil}
	node2 := &linkedlist.Node{Data: 2, Next: nil}
	node3 := &linkedlist.Node{Data: 3, Next: nil}
	node4 := &linkedlist.Node{Data: 4, Next: nil}
	node5 := &linkedlist.Node{Data: 5, Next: nil}
	node6 := &linkedlist.Node{Data: 6, Next: nil}
	node7 := &linkedlist.Node{Data: 7, Next: nil}
	node8 := &linkedlist.Node{Data: 8, Next: nil}
	node9 := &linkedlist.Node{Data: 9, Next: nil}

	list := linkedlist.List{node1}
	linkedlist.InsertAfter(node1, node2)
	linkedlist.InsertAfter(node2, node3)
	linkedlist.InsertAfter(node3, node4)
	linkedlist.InsertAfter(node4, node5)
	linkedlist.InsertAfter(node5, node6)
	linkedlist.InsertAfter(node6, node7)
	linkedlist.InsertAfter(node7, node8)
	linkedlist.InsertAfter(node8, node9)

	return list
}

func TestSimple(t *testing.T) {
	kth := 5
	answer := Simple(linkedListCase1(), kth)
	if answer.Data.(int) != 10-kth {
		t.Errorf("Expected %d, but got %d", 10-kth, answer.Data.(int))
	}
}

/*
func TestRecursion(t *testing.T) {
	list := linkedListCase1()
	kth := 1
	var answer *linkedlist.Node
	Recursion(list.FirstNode, 0)

	if answer.Data.(int) != 10-kth {
		t.Errorf("Expected %d, but got %d", 10-kth, answer.Data.(int))
	}
}
*/
