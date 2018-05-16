package question1

import (
	"testing"

	"github.com/crazcalm/interview-prep/data_structures/linked_list"
)

func linkedListCase1() (linkedlist.List, []int) {
	node1 := &linkedlist.Node{Data: 1, Next: nil}
	node2 := &linkedlist.Node{Data: 1, Next: nil}
	node3 := &linkedlist.Node{Data: 1, Next: nil}
	node4 := &linkedlist.Node{Data: 2, Next: nil}
	node5 := &linkedlist.Node{Data: 3, Next: nil}
	node6 := &linkedlist.Node{Data: 4, Next: nil}
	node7 := &linkedlist.Node{Data: 2, Next: nil}
	node8 := &linkedlist.Node{Data: 3, Next: nil}
	node9 := &linkedlist.Node{Data: 3, Next: nil}

	list := linkedlist.List{node1}
	linkedlist.InsertAfter(node1, node2)
	linkedlist.InsertAfter(node2, node3)
	linkedlist.InsertAfter(node3, node4)
	linkedlist.InsertAfter(node4, node5)
	linkedlist.InsertAfter(node5, node6)
	linkedlist.InsertAfter(node6, node7)
	linkedlist.InsertAfter(node7, node8)
	linkedlist.InsertAfter(node8, node9)

	return list, []int{1, 2, 3, 4}
}

func TestHash(t *testing.T) {
	//test1
	testCase, solution := linkedListCase1()
	answerList := Hash(testCase)

	var count int
	for n := answerList.FirstNode; n != nil; n = n.Next {
		value := n.Data.(int)
		if value != solution[count] {
			t.Errorf("Expected %d at node %d, but got %d", solution[count], count, value)
		}
		count++
	}

	//Test 2
	node27 := &linkedlist.Node{Data: 1, Next: nil}
	list := linkedlist.List{node27}
	answerList = Hash(list)
	solution = []int{1}

	count = 0
	for n := answerList.FirstNode; n != nil; n = n.Next {
		value := n.Data.(int)
		if value != solution[count] {
			t.Errorf("Expected %d at node %d, but got %d", solution[count], count, value)
		}
		count++
	}

}

func TestFollowUp(t *testing.T) {
	//test1
	testCase, solution := linkedListCase1()
	answerList := FollowUp(testCase)

	var count int
	for n := answerList.FirstNode; n != nil; n = n.Next {
		value := n.Data.(int)
		if value != solution[count] {
			t.Errorf("Expected %d at node %d, but got %d", solution[count], count, value)
		}
		count++
	}

	//Test 2
	node7 := &linkedlist.Node{Data: 1, Next: nil}
	list := linkedlist.List{node7}
	answerList = FollowUp(list)
	solution = []int{1}

	count = 0
	for n := answerList.FirstNode; n != nil; n = n.Next {
		value := n.Data.(int)
		if value != solution[count] {
			t.Errorf("Expected %d at node %d, but got %d", solution[count], count, value)
		}
		count++
	}

}
