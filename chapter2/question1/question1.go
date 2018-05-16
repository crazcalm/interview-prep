package question1

import (
	"github.com/crazcalm/interview-prep/data_structures/linked_list"
)

/*
Write code to remove duplicates from an unsorted linked list.

Follow up: How would you solve this problem if a temporary buffer was not allowed.
*/

//FollowUp -- follow up answer
func FollowUp(list linkedlist.List) linkedlist.List {
	currentNode := list.FirstNode
	for currentNode != nil {
		// Remove all future nodes that have the same value
		runner := currentNode
		for runner.Next != nil {
			if runner.Next.Data.(int) == currentNode.Data.(int) {
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}
		currentNode = currentNode.Next
	}
	return list
}

//Hash -- Uses a Hash in the solution. This solution assumes that was
//are only dealing with integers
func Hash(list linkedlist.List) linkedlist.List {
	/*
		Notes: I decided to copy each unique node into a new linked list because
			   I am afraid of modifying the iteration path while iterating over
			   the list.
	*/
	var counter = make(map[int]int)
	var result linkedlist.List
	var currentNode *linkedlist.Node
	for n := list.FirstNode; n != nil; n = n.Next {
		value := n.Data.(int)
		_, ok := counter[value]
		if !ok { //Case: data is unique

			//Add it to hash
			counter[value] = 1

			//Copy the node data into a new node
			newNode := &linkedlist.Node{Data: n.Data, Next: nil}

			//Add it to the list
			if result.FirstNode == nil {
				result = linkedlist.InsertBeginning(result, newNode)
			} else {
				linkedlist.InsertAfter(currentNode, newNode)
			}
			//Set current node
			currentNode = n

		}
	}
	return result
}
