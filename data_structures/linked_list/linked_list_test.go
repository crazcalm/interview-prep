package linkedlist

import (
	"strings"
	"testing"
)

func TestRemoveBeginning(t *testing.T) {
	node1 := &Node{Data: "Being Removed", Next: nil}
	node2 := &Node{Data: "Keeping Node", Next: nil}
	list := List{node1}
	InsertAfter(node1, node2)

	//Test
	list = RemoveBeginning(list)

	if list.FirstNode != node2 {
		t.Errorf("Expected list.FirstNode to be %v, but got %v", node2, list.FirstNode)
	}
}

func TestRemoveAfter(t *testing.T) {
	node1 := &Node{Data: "Keeing Node", Next: nil}
	node2 := &Node{Data: "Being Removed", Next: nil}
	list := List{node1}
	expectedNumOfNodes := 1

	InsertAfter(node1, node2)

	//Actual test
	RemoveAfter(node1)

	var count int
	for n := list.FirstNode; n != nil; n = n.Next {
		count++
	}

	if count != expectedNumOfNodes {
		t.Errorf("Expected %d node(s), but got %d", expectedNumOfNodes, count)
	}

}

func TestInsertBeginning(t *testing.T) {
	node := &Node{Data: "OldStart", Next: nil}
	list := List{node}
	newNode := &Node{Data: "New Start", Next: nil}

	list = InsertBeginning(list, newNode)

	//Test -- The only needed check
	if newNode != list.FirstNode {
		t.Errorf("Expected list.FirstNode to be %v, but got %v: list -- %v", newNode, list.FirstNode, list)
	}

	if !strings.EqualFold(newNode.Data.(string), list.FirstNode.Data.(string)) {
		t.Errorf("Expected the data of the first node to be %s, but got %s", newNode.Data.(string), list.FirstNode.Data.(string))
	}

}

func TestInsertAfter(t *testing.T) {
	node := &Node{Data: 19, Next: nil}
	list := List{node}
	newNode := &Node{Data: "5", Next: nil}
	expectedNumOfNodes := 2

	InsertAfter(node, newNode)

	//Test
	var count int
	for n := list.FirstNode; n != nil; n = n.Next {
		count++
	}

	if count != expectedNumOfNodes {
		t.Errorf("Expected list to have %d node, but got %d. List = %v", expectedNumOfNodes, count, list)
	}
}
