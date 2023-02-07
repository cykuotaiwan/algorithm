package redBlackTree

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
	for node != tree.Root && node.color == BLACK {
		if node == node.siblingLeft() {
			// case 1
			//     node is the left child
			sibling := node.siblingRight()
			if sibling.color == RED {
				sibling.color = BLACK
				node.Parent.color = RED
				tree.rotateLeft(node.Parent)
				sibling = node.Parent.Right
			}

			if sibling.Left.color == BLACK && sibling.Right.color == BLACK {
				// case 2
				//     node's sibling is black
				//     and both of the sibling's children are also black
				sibling.color = RED
				node = node.Parent
			} else if sibling.Right.color == BLACK {
				// case 3
				//     node's sibling is black
				//     and sibling's left child is red, right child is black
				sibling.Left.color = BLACK
				sibling.color = RED
				tree.rotateRight(sibling)
				sibling = node.Parent.Right
			}
			// case 4
			//     node's sibling is black
			//     and sibling's right child is red
			sibling.color = node.Parent.color
			node.Parent.color = BLACK
			sibling.Right.color = BLACK
			tree.rotateLeft(node.Parent)
			node = tree.Root
		} else {
			if node == node.siblingRight() {
				// case 1
				//     node is the right child
				sibling := node.siblingLeft()
				if sibling.color == RED {
					sibling.color = BLACK
					node.Parent.color = RED
					tree.rotateRight(node.Parent)
					sibling = node.Parent.Left
				}

				if sibling.Right.color == BLACK && sibling.Left.color == BLACK {
					// case 2
					//     node's sibling is black
					//     and both of the children are also black
					sibling.color = RED
					node = node.Parent
				} else if sibling.Left.color == BLACK {
					// case 3
					//     node's sibling is black
					//     and sibling's left child is red, right child is black
					sibling.Right.color = BLACK
					sibling.color = RED
					tree.rotateLeft(sibling)
					sibling = node.Parent.Left
				}

				// case 4
				//     node's sibling is black
				//     and sibling's right child is red
				sibling.color = node.Parent.color
				node.Parent.color = BLACK
				sibling.Left.color = BLACK
				tree.rotateRight(node.Parent)
				node = tree.Root
			}
		}
	}
	node.color = BLACK
}
