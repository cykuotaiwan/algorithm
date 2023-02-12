package redBlackTree

func (tree *RedBlackTree) replaceNode(old *TreeNode, new *TreeNode) {
	if old.Parent == nil {
		tree.Root = new
		return
	}

	if old == old.Parent.Left { // node is left child
		old.Parent.Left = new
	} else { // node is right child
		old.Parent.Right = new
	}

	if new != nil {
		new.Parent = old.Parent
	}
}

func (tree *RedBlackTree) rotateLeft(node *TreeNode) {
	right := node.Right
	tree.replaceNode(node, right)
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Left = node
	node.Parent = right
}

func (tree *RedBlackTree) rotateRight(node *TreeNode) {
	left := node.Left
	tree.replaceNode(node, left)
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Right = node
	node.Parent = left
}
