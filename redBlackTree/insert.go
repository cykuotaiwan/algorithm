package redBlackTree

import (
	"fmt"
)

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

		if node == nil {
			node = tmp
			break
		}
	}
	fmt.Printf("%p ; %v \t<-- parent\n", node, node)
	insertNode.Parent = node
	if tmp == nil { // tree.Root is nil
		tree.Root = insertNode
	} else if key < tmp.Key {
		tmp.Left = insertNode
	} else {
		tmp.Right = insertNode
	}
	fmt.Printf("%p ; %v \t<-- insert\n", insertNode, insertNode)
	tree.Visualize("prefix")

	tree.insertFix(insertNode)

	tree.Visualize("postfix")
	fmt.Printf("%p ; %v \t<-- parent fixed\n\n", node, node)
}

func (tree *RedBlackTree) insertFix(node *TreeNode) {
	if node.Parent == nil { // node is tree.Root
		node.color = BLACK
		return
	}

	if node.color == BLACK {
		return
	}

	fmt.Printf("%p ; %v \t<-- fix before\n", node, node)

	for node.Parent.color == RED {
		if node.Parent == node.grandparent().Left {
			uncle := node.uncleRight()
			if uncle.color == RED {
				// case 1
				//     node's uncle is red
				fmt.Println("[Case Left 1]")

				node.Parent.color = BLACK
				uncle.color = BLACK

				node.grandparent().color = RED
				node = node.grandparent()
				fmt.Printf("%p ; %v \t<-- fix after case 1\n", node, node)
			} else if node == node.Parent.Right {
				// case 2
				//     node's uncle is black and node is a right child
				fmt.Println("[Case Left 2]")

				node = node.Parent
				tree.rotateLeft(node)
			}
			// case 3
			//     node's uncle is black and node is a left child
			fmt.Println("[Case Left 3]")
			node.Parent.color = BLACK
			node.grandparent().color = RED
			tree.rotateRight(node.grandparent())
		} else {
			if node.Parent == node.grandparent().Right {
				uncle := node.uncleLeft()
				if uncle.color == RED {
					// case 1
					//     node's uncle is red
					fmt.Println("[Case Right 1]")
					node.Parent.color = BLACK
					uncle.color = BLACK

					node.grandparent().color = RED
					node = node.grandparent()
				}
			} else if node == node.Parent.Left {
				// case 2
				//     node's uncle is black and node is a right child
				fmt.Println("[Case Right 2]")
				node = node.Parent
				tree.rotateRight(node)
			}
			// case 3
			//     node's uncle is black and node is a left child
			fmt.Println("[Case Right 3]")
			node.Parent.color = BLACK
			node.grandparent().color = RED
			tree.rotateLeft(node.grandparent())
		}
	}
	fmt.Printf("%p ; %v \t<-- fix after\n", node, node)

	tree.Root.color = BLACK
}
