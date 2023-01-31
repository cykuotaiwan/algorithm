package binarySearchTree

func (tree *BinarySearchTree) MinNode() *TreeNode {
	return minNode(tree.root)
}

func minNode(node *TreeNode) *TreeNode {
	tmp := node
	for tmp.left != nil {
		tmp = tmp.left
	}
	return tmp
}

func (tree *BinarySearchTree) MaxNode() *TreeNode {
	return maxNode(tree.root)
}

func maxNode(node *TreeNode) *TreeNode {
	tmp := node
	for tmp.right != nil {
		tmp = tmp.right
	}
	return tmp
}
