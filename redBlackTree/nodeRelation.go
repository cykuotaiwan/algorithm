package redblacktree

func (node *TreeNode) siblingRight() *TreeNode {
	if node == nil || node.Parent == nil {
		return nil
	}
	return node.Parent.Right
}

func (node *TreeNode) siblingLeft() *TreeNode {
	if node == nil || node.Parent == nil {
		return nil
	}
	return node.Parent.Left
}

func (node *TreeNode) uncleRight() *TreeNode {
	if node == nil || node.Parent == nil || node.Parent.Parent == nil {
		return nil
	}
	return node.Parent.siblingRight()
}

func (node *TreeNode) uncleLeft() *TreeNode {
	if node == nil || node.Parent == nil || node.Parent.Parent == nil {
		return nil
	}
	return node.Parent.siblingLeft()
}

func (node *TreeNode) grandparent() *TreeNode {
	if node == nil || node.Parent == nil {
		return nil
	}
	return node.Parent.Parent
}
