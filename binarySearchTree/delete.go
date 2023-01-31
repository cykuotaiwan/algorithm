package binarySearchTree

func (tree *BinarySearchTree) DeleteElement(val int) {
	deleteNode(tree.root, val)
}

func deleteNode(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}
	if val < node.val {
		node.left = deleteNode(node.left, val)
	}
	if val > node.val {
		node.right = deleteNode(node.right, val)
	}

	// val == node.val
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}
	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}

	leftMostRightNode := node.right
	for {
		if leftMostRightNode != nil && leftMostRightNode.left != nil {
			leftMostRightNode = leftMostRightNode.left
		} else {
			break
		}
	}
	node.val = leftMostRightNode.val
	node.right = deleteNode(node.right, node.val)
	return node
}
