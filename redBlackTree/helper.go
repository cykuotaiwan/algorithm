package redBlackTree

func getColor(node *TreeNode) color {
	if node == nil {
		return BLACK
	}
	return node.color
}
