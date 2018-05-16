package question2

import (
	"github.com/crazcalm/interview-prep/data_structures/linked_list"
)

/*
Question: Return Kth to Last: Implement an algorithm to find the kth to last element of a singly
          linked list
*/

func getElement(node *linkedlist.Node, count, wanted int) *linkedlist.Node {
	if node == nil {
		return nil
	}

	if count == wanted-1 { //incrimented down because the algrithm starts the count at 0
		return node
	}
	count++
	return getElement(node.Next, count, wanted)
}

func lengthRecursion(node *linkedlist.Node) int {
	if node == nil {
		return 0
	}
	return lengthRecursion(node.Next) + 1
}

//Recursion -- This solution uses recursion
func Recursion(list linkedlist.List, kth int) *linkedlist.Node {
	length := lengthRecursion(list.FirstNode)
	wanted := length - kth
	count := 0 //we start counting at 0
	return getElement(list.FirstNode, count, wanted)
}

//Simple -- This solution uses one loop to find out the lengh of the list
//and another loop to return the kth Node
func Simple(list linkedlist.List, kth int) (result *linkedlist.Node) {
	var length int
	for n := list.FirstNode; n != nil; n = n.Next {
		length++
	}

	if length <= kth {
		return
	}

	wantedIndex := length - kth - 1 //incriments down by 1 because the algorith starts count at 0
	var count int
	for n := list.FirstNode; n != nil; n = n.Next {
		if count == wantedIndex {
			result = n
			break
		}
		count++
	}

	return
}
