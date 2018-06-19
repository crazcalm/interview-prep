package bst

//Element -- Element of a binary tree
type Element struct {
	Key                 uint32
	Value               interface{}
	Parent, Left, Right *Element
}

//BinarySearchTree -- BST
type BinarySearchTree struct {
	Root *Element
}

//NewTree -- returns a bst tree
func NewTree() *BinarySearchTree {
	return new(BinarySearchTree)
}

//RecursiveSearch --
func RecursiveSearch(cur *Element, key uint32) *Element {
	//Base Case: Did not find anything
	if cur == nil {
		return nil
	}

	//Base Case: Found it!
	if cur.Key == key {
		return cur
	}

	//Recursive cases
	if key < cur.Key {
		return RecursiveSearch(cur.Left, key)
	}
	return RecursiveSearch(cur.Right, key)
}

//Search -- Searches for element by key
func (t *BinarySearchTree) Search(key uint32) *Element {
	for cur := t.Root; cur != nil; {
		if cur.Key == key {
			return cur
		} else if key < cur.Key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return nil
}

//RecursiveInsert --
func RecursiveInsert(cur, node *Element, tree *BinarySearchTree) *BinarySearchTree {
	//Base Case: Root is nil. Need to Create a new tree
	if cur == nil {
		cur = node
		return &BinarySearchTree{Root: &Element{Value: node.Value, Key: node.Key, Parent: nil, Left: nil, Right: nil}}
	}

	if cur.Left == nil && node.Key < cur.Key {
		cur.Left = node
		node.Parent = cur
		return tree
	}

	if cur.Right == nil && node.Key >= cur.Key {
		cur.Right = node
		node.Parent = cur
		return tree
	}

	//Recursive cases
	if node.Key < cur.Key {
		return RecursiveInsert(cur.Left, node, tree)
	}
	return RecursiveInsert(cur.Right, node, tree)

}

//Insert -- adds a new element to the tree
func (t *BinarySearchTree) Insert(n *Element) *Element {
	var target *Element

	//Find suitable target node to unload this new node to
	for cur := t.Root; cur != nil; {
		target = cur
		if n.Key < cur.Key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}

	n.Parent = target
	if target == nil {
		t.Root = n
	} else if n.Key < target.Key {
		target.Left = n
	} else {
		target.Right = n
	}
	return n
}

//Delete -- removes an item from the tree
func (t *BinarySearchTree) Delete(key uint32) *Element {
	deleteNonCompleteNode := func(node *Element) {
		var reConnectedNode *Element
		if node.Left == nil {
			reConnectedNode = node.Right
		} else {
			reConnectedNode = node.Left
		}

		if reConnectedNode != nil {
			reConnectedNode.Parent = node.Parent
		}

		if node.Parent == nil {
			t.Root = reConnectedNode
		} else if node.Parent.Right == node {
			node.Parent.Right = reConnectedNode
		} else {
			node.Parent.Left = reConnectedNode
		}
		node = nil
	}

	node := t.Search(key)
	if node == nil { //no Node to delete
		return node
	}

	if node.Left == nil || node.Right == nil {
		deleteNonCompleteNode(node)
	} else {
		successor := t.Successor(node, t.Root)
		_key, _value := successor.Key, successor.Value
		deleteNonCompleteNode(successor)
		node.Key, node.Value = _key, _value
	}
	return node
}

//RecursiveMin --
func RecursiveMin(cur *Element) *Element {
	if cur == nil {
		return nil
	}

	if cur.Left == nil {
		return cur
	}

	return RecursiveMin(cur.Left)
}

//Min -- Returns the element with the smallest key
func (t *BinarySearchTree) Min(cur *Element) *Element {
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur
}

//Successor -- Used to help with deletion
func (t *BinarySearchTree) Successor(n, root *Element) *Element {
	if n == nil {
		return nil
	}
	if n.Right != nil {
		return t.Min(n.Right)
	}
	cur := n
	for cur.Parent != nil && cur.Parent.Left != cur {
		cur = cur.Parent
	}
	return cur.Parent

}

//TraverseInOrder -- Traverses the tree and calls the callback
func TraverseInOrder(cur *Element, callback func(cur *Element)) {
	if cur == nil {
		return
	}

	TraverseInOrder(cur.Left, callback)
	callback(cur)
	TraverseInOrder(cur.Right, callback)
}

//TraverseReverseOrder -- Traverses the tree and calls the callback
func TraverseReverseOrder(cur *Element, callback func(cur *Element)) {
	if cur == nil {
		return
	}

	TraverseReverseOrder(cur.Right, callback)
	callback(cur)
	TraverseReverseOrder(cur.Left, callback)
}
