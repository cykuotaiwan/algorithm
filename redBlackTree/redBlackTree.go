package redBlackTree

type color bool

const (
	BLACK, RED color = true, false
)

type TreeNode struct {
	Val    int
	Key    int
	color  color
	Parent *TreeNode
	Left   *TreeNode
	Right  *TreeNode
}

type RedBlackTree struct {
	Root *TreeNode
}

/*
Red Black Tree Properties
1. Every node is either red or black
2. The root is black
3. Every leaf (nil) is black
4. If a node is red, both of its children are black
	(aka no consecutive 2 red nodes)
5. For each node, all paths from the node to the leaf (nil)
	passes through exact the same number of black nodes
*/
