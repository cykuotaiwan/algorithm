package redblacktree

func (tree *RedBlackTree) search(key int) *TreeNode {
	node := tree.Root
	for node != nil {
		if node.Key == key {
			return node
		} else if node.Key < key {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return nil
}
