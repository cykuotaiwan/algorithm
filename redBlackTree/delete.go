package redBlackTree

func (tree *RedBlackTree) Delete(key int) {
	var child *TreeNode
	node := tree.search(key)
	if node == nil {
		return
	}
	colorOriginal := node.color

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
		tree.deleteCase1(child)
	}
}

func (tree *RedBlackTree) deleteCase1(node *TreeNode) {
	if node.Parent == nil {
		return
	}
	tree.deleteCase2(node)
}

func (tree *RedBlackTree) deleteCase2(node *TreeNode) {
	sibling := node.sibling()

	if getColor(sibling) == RED {
		node.Parent.color = RED
		sibling.color = BLACK
		if node == node.sibling() {
			tree.rotateLeft(node.Parent)
		} else {
			tree.rotateRight(node.Parent)
		}
	}

	tree.deleteCase3(node)
}

func (tree *RedBlackTree) deleteCase3(node *TreeNode) {
	sibling := node.sibling()

	if getColor(node.Parent) == BLACK &&
		getColor(sibling) == BLACK &&
		getColor(sibling.Left) == BLACK &&
		getColor(sibling.Right) == BLACK {
		sibling.color = RED
		tree.deleteCase1(node.Parent)

	} else {
		tree.deleteCase4(node)
	}
}

func (tree *RedBlackTree) deleteCase4(node *TreeNode) {
	sibling := node.sibling()

	if getColor(node.Parent) == RED &&
		getColor(sibling) == BLACK &&
		getColor(sibling.Left) == BLACK &&
		getColor(sibling.Right) == BLACK {
		sibling.color = RED
		node.Parent.color = BLACK

	} else {
		tree.deleteCase5(node)
	}
}

func (tree *RedBlackTree) deleteCase5(node *TreeNode) {
	sibling := node.sibling()

	if node == node.Parent.Left &&
		getColor(sibling) == BLACK &&
		getColor(sibling.Left) == RED &&
		getColor(sibling.Right) == BLACK {
		sibling.color = RED
		sibling.Left.color = BLACK
		tree.rotateRight(sibling)
	} else if node == node.Parent.Right &&
		getColor(sibling) == BLACK &&
		getColor(sibling.Right) == RED &&
		getColor(sibling.Left) == BLACK {
		sibling.color = RED
		sibling.Right.color = BLACK
		tree.rotateLeft(sibling)
	}

	tree.deleteCase6(node)
}

func (tree *RedBlackTree) deleteCase6(node *TreeNode) {
	sibling := node.sibling()

	sibling.color = getColor(node.Parent)
	node.Parent.color = BLACK

	if node == node.siblingLeft() && getColor(sibling.Right) == RED {
		sibling.Right.color = BLACK
		tree.rotateLeft(node.Parent)
	} else if getColor(sibling.Left) == RED {
		sibling.Left.color = BLACK
		tree.rotateRight(node.Parent)
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
