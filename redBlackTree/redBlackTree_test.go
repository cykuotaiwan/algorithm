package redBlackTree_test

import (
	rbt "algorithms/redBlackTree"
	"testing"
)

func TestRedBlackTree(t *testing.T) {
	var tree rbt.RedBlackTree
	tree.Insert(9, 9)
	tree.Insert(8, 8)
	tree.Insert(16, 16)
	tree.Insert(7, 7)

	tree.Visualize("testMain")
}
