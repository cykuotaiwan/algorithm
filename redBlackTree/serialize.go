package redBlackTree

import "fmt"

func (node *TreeNode) Print(nodeName string) {
	fmt.Printf("%s: %p, %v\n", nodeName, node, node)
}
