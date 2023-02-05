package redblacktree

func (tree *RedBlackTree) transplant(old *TreeNode, new *TreeNode) {
	if old.Parent == nil { // old is root
		tree.Root = new
	} else if old == old.siblingLeft() {
		old.Parent.Left = new
	} else {
		old.Parent.Right = new
	}

	new.Parent = old.Parent
}
