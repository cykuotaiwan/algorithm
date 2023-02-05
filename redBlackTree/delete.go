package redblacktree

func (tree *RedBlackTree) Delete(key int) {
	var child *TreeNode
	node := tree.search(key)
	colorOriginal := node.color
	if node == nil {
		return
	}

	if node.Left == nil {
		child = node.Right
		tree.transplant(node, child)
	} else if node.Right == nil {
		child = node.Left
		tree.transplant(node, child)
	} else {
		tmp := node.Right.minNode()
		colorOriginal = tmp.color
		child = node.Right
		if tmp.Parent == node {
			child.Parent = tmp
		} else {
			tree.transplant(tmp, tmp.Right)
			tmp.Right = child.Right
			tmp.Right.Parent = tmp
		}

		tree.transplant(node, tmp)
		tmp.Left = node.Left
		tmp.Left.Parent = tmp
		tmp.color = node.color
	}

	if colorOriginal == BLACK {
		tree.deleteFix(child)
	}
}

func (tree *RedBlackTree) deleteFix(node *TreeNode) {

}
