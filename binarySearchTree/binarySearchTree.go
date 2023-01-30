package binarySearchTree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTree(arr []int) *TreeNode {
	var root *TreeNode

	for _, val := range arr {
		node := &TreeNode{val, nil, nil}
		fmt.Println(node)
	}

	return root
}
