package redblacktree

// assume no duplicate keys
func (tree *RedBlackTree) Insert(key int, val int) {
	insertNode := &TreeNode{
		Key:   key,
		Val:   val,
		color: RED,
		Left:  nil,
		Right: nil,
	}

	var tmp *TreeNode
	node := tree.Root

	for node != nil {
		tmp = node
		if key < tmp.Key {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	insertNode.Parent = node
	if tmp == nil { // tree.Root is nil
		tree.Root = insertNode
	} else if key < tmp.Key {
		tmp.Left = insertNode
	} else {
		tmp.Right = insertNode
	}
	tree.insertFix(insertNode)
}

func (tree *RedBlackTree) insertFix(node *TreeNode) {
	if node.Parent == nil { // node is tree.Root
		node.color = BLACK
		return
	}

	if node.color == BLACK {
		return
	}

	for node.Parent.color == RED {
		if node.Parent == node.grandparent().Left {
			uncle := node.uncleRight()
			if uncle.color == RED {
				// case 1
				//     node's uncle is red
				node.Parent.color = BLACK
				uncle.color = BLACK

				node.grandparent().color = RED
				node = node.grandparent()
			} else if node == node.Parent.Right {
				// case 2
				//     node's uncle is black and node is a right child
				node = node.Parent
				tree.rotateLeft(node)
			}
			// case 3
			//     node's uncle is black and node is a left child
			node.Parent.color = BLACK
			node.grandparent().color = RED
			tree.rotateRight(node.grandparent())
		} else {
			if node.Parent == node.grandparent().Right {
				uncle := node.uncleLeft()
				if uncle.color == RED {
					// case 1
					//     node's uncle is red
					node.Parent.color = BLACK
					uncle.color = BLACK

					node.grandparent().color = RED
					node = node.grandparent()
				}
			} else if node == node.Parent.Left {
				// case 2
				//     node's uncle is black and node is a right child
				node = node.Parent
				tree.rotateRight(node)
			}
			// case 3
			//     node's uncle is black and node is a left child
			node.Parent.color = BLACK
			node.grandparent().color = RED
			tree.rotateLeft(node.grandparent())
		}
	}

	tree.Root.color = BLACK
}

func insertCase1(uncle *TreeNode, node *TreeNode) {

}
