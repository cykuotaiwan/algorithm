package binarySearchTree

func (tree *BinarySearchTree) Search(val int) *TreeNode {
	return search(tree.root, val)
}

func search(node *TreeNode, val int) *TreeNode {
	tmp := node
	for tmp != nil && val != tmp.val {
		if val < tmp.val {
			tmp = tmp.left
		} else {
			tmp = tmp.right
		}
	}

	return tmp
}
