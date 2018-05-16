package question2

import (
	"github.com/crazcalm/interview-prep/data_structures/linked_list"
)

/*
Question: Return Kth to Last: Implement an algorithm to find the kth to last element of a singly
          linked list
*/

/*
func Recursion(node *linkedlist.Node, kth, count int, answer *linkedlist.Node) {
	if node == nil {
		return
	}
	count++
	Recursion(node.Next, kth, count, answer)
	fmt.Printf("kth to last: count - kth = %d\n", count-kth)

	fmt.Printf("Count: %d -- Node: %d\n", count, node.Data.(int))
}

*/

/*
//Recursion -- This solution uses recursion
func Recursion(node *linkedlist.Node, count int) {
	if node == nil {
		return
	}
	fmt.Printf("Count %d -- Node %v\n", count, node)
	count++
	Recursion(node.Next, count)

	fmt.Printf("Count %d -- Node %v\n", count, node)
}
*/

//Simple -- This solution uses one loop to find out the lengh of the list
//and another loop to return the kth Node
func Simple(list linkedlist.List, kth int) (result *linkedlist.Node) {
	var length int
	for n := list.FirstNode; n != nil; n = n.Next {
		length++
	}

	if length < kth {
		return
	}

	wantedIndex := length - kth
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
