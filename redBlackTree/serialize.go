package redBlackTree

import "fmt"

func (node *TreeNode) print(nodeName string) {
	fmt.Printf("%s: %p, %v\n", nodeName, node, node)
}
