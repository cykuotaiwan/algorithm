package binarySearchTree

func (tree *BinarySearchTree) Inorder() []int {
	queue := make([]int, 0)

	queue = inorder(tree.root, queue)

	return queue
}

func inorder(node *TreeNode, queue []int) []int {
	if node != nil {
		queue = inorder(node.left, queue)
		queue = append(queue, node.val)
		queue = inorder(node.right, queue)
	}
	return queue
}

func (tree *BinarySearchTree) Preorder() []int {
	var queue []int

	queue = preorder(tree.root, queue)

	return queue
}

func preorder(node *TreeNode, queue []int) []int {
	if node != nil {
		queue = append(queue, node.val)
		queue = preorder(node.left, queue)
		queue = preorder(node.right, queue)
	}

	return queue
}

func (tree *BinarySearchTree) Postorder() []int {
	var queue []int

	queue = postorder(tree.root, queue)

	return queue
}

func postorder(node *TreeNode, queue []int) []int {
	if node != nil {
		queue = postorder(node.left, queue)
		queue = postorder(node.right, queue)
		queue = append(queue, node.val)
	}

	return queue
}
