package binarySearchTree_test

import (
	"algorithms/binarySearchTree"
	"algorithms/util"
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	arr := util.GenerateRandArr(5, true)
	tree := new(binarySearchTree.BinarySearchTree)

	fmt.Println(arr)

	for _, val := range arr {
		tree.InsertElement(val)
	}

	fmt.Println(tree.Inorder())
	fmt.Println(tree.Preorder())
	fmt.Println(tree.Postorder())
}
