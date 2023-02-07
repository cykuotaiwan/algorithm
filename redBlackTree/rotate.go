package redBlackTree

func (tree *RedBlackTree) rotateLeft(node *TreeNode) {
	right := node.Right

	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}

	right.Parent = node.Parent
	if node.Parent != nil { // node is root
		tree.Root = right
	} else if node == node.Parent.Left { // node is left child
		node.Parent.Left = right
	} else { // node is right child
		node.Parent.Right = right
	}

	right.Left = node
	node.Parent = right
}

func (tree *RedBlackTree) rotateRight(node *TreeNode) {
	left := node.Left

	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}

	left.Parent = node.Parent
	if node.Parent != nil { // node is root
		tree.Root = left
	} else if node == node.Parent.Left { // node is left child
		node.Parent.Left = left
	} else { // node is right child
		node.Parent.Right = left
	}

	left.Right = node
	node.Parent = left
}
