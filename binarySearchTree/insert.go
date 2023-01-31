package binarySearchTree

func (tree *BinarySearchTree) InsertElement(value int) {
	newNode := &(TreeNode{val: value})

	if tree.root == nil {
		tree.root = newNode
	} else {
		insertNode(tree.root, newNode)
	}
	// fmt.Println("Tree", tree)
	// fmt.Printf("  TreeRoot %v %p\n", tree.root, tree.root)
}

func insertNode(root *TreeNode, newNode *TreeNode) {
	if newNode.val < root.val {
		if root.left == nil {
			root.left = newNode
		} else {
			insertNode(root.left, newNode)
		}
	}

	if newNode.val > root.val {
		if root.right == nil {
			root.right = newNode
		} else {
			insertNode(root.right, newNode)
		}
	}

	// fmt.Printf("    Root %v %p\n", root, root)
}
