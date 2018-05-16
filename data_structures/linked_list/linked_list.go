package linkedlist

//Node -- A node in a linked list
type Node struct {
	Data interface{}
	Next *Node
}

//List -- Lists points to the first node in a linked list
type List struct {
	FirstNode *Node
}

//InsertAfter -- Inserts a Node after a specific node
func InsertAfter(node, newNode *Node) {
	newNode.Next = node.Next
	node.Next = newNode
}

//InsertBeginning -- Inserts a new first node into the list
func InsertBeginning(list List, newNode *Node) List {
	newNode.Next = list.FirstNode
	list.FirstNode = newNode
	return list
}

//RemoveAfter -- Removes a Node located after a specific Node.
//Note: The removed node is being unreferenced.
//The garbage collector should pick it up.
func RemoveAfter(node *Node) {
	node.Next = node.Next.Next
}

//RemoveBeginning -- Removes the first Node in the List.
//Note: The removed not is being unreferenced.
//The garbage collector should pick it up.
//Note: list.FirstNode will be set to nil when removing the last node in the list.
func RemoveBeginning(list List) List {
	list.FirstNode = list.FirstNode.Next
	return list
}
