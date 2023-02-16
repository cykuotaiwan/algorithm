package redBlackTree

func (node *TreeNode) minNode() *TreeNode {
	curr := node
	for curr.Left != nil {
		curr = curr.Left
	}
	return curr
}

func (node *TreeNode) maxNode() *TreeNode {
	curr := node
	for curr.Right != nil {
		curr = curr.Right
	}
	return curr
}
