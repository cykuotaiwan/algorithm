package redBlackTree

import "fmt"

func (tree *RedBlackTree) Insert(key int, val int) {
	var node *TreeNode

	if tree.Root == nil {
		node = &TreeNode{Key: key, Val: val, color: RED}
		tree.Root = node
	} else {
		curr := tree.Root
		flag := true
		for flag {
			if curr.Key == key {
				curr.Key = key
				curr.Val = val
			} else if curr.Key > key {
				if curr.Left == nil {
					node = &TreeNode{Key: key, Val: val, color: RED}
					node.Parent = curr
					curr.Left = node
					flag = false
				} else {
					curr = curr.Left
				}
			} else { // curr.Key > key
				if curr.Right == nil {
					node = &TreeNode{Key: key, Val: val, color: RED}
					node.Parent = curr
					curr.Right = node
					flag = false
				} else {
					curr = curr.Right
				}
			}
		}
	}
	tree.Visualize("prefix")
	tree.insertCase1(node)
	fmt.Println()
	tree.Visualize("postfix")
}

// Reference: https://en.wikipedia.org/wiki/Red%E2%80%93black_tree

func (tree *RedBlackTree) insertCase1(node *TreeNode) {
	node.print("case1 node")
	if node.Parent == nil { // node is root
		node.color = BLACK
		tree.Visualize("case1")
	} else {
		tree.insertCase2(node)
	}
}

func (tree *RedBlackTree) insertCase2(node *TreeNode) {
	node.print("case2 node")
	if getColor(node) == BLACK {
		tree.Visualize("case2")
		return
	} else {
		tree.insertCase3(node)
	}
}

func (tree *RedBlackTree) insertCase3(node *TreeNode) {
	node.print("case3 node")
	uncle := node.uncle()
	uncle.print("\tuncle")
	tree.Visualize("precase3")

	if getColor(uncle) == RED {
		node.Parent.color = BLACK
		uncle.color = BLACK
		node.grandparent().color = RED
		tree.insertCase1(node.grandparent())
		tree.Visualize("case3")
	} else {
		tree.insertCase4(node)
	}
}

func (tree *RedBlackTree) insertCase4(node *TreeNode) {
	node.print("case4 node")
	grandparent := node.grandparent()
	grandparent.print("\tgrandparent")

	if grandparent == nil {
		return
	}

	if node == node.Parent.Right && node.Parent == grandparent.Left {
		tree.rotateLeft(node.Parent)
		node = node.Left
	} else if node == node.Parent.Left && node.Parent == grandparent.Right {
		tree.rotateRight(node.Parent)
		node = node.Right
	}
	tree.Visualize("case4")

	tree.insertCase5(node)
}

func (tree *RedBlackTree) insertCase5(node *TreeNode) {
	node.print("case5 node")
	node.Parent.color = BLACK
	grandparent := node.grandparent()
	grandparent.print("\tgrandparent")
	grandparent.color = RED

	if node == node.Parent.Left && node.Parent == grandparent.Left {
		tree.rotateRight(grandparent)
	} else if node == node.Parent.Right && node.Parent == grandparent.Right {
		tree.rotateLeft(grandparent)
	}
	tree.Visualize("case5")
}
