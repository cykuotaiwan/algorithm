package redBlackTree_test

import (
	rbt "algorithms/redBlackTree"
	"testing"
)

func TestRedBlackTreeInsert(t *testing.T) {
	tree := rbt.NewRBTree()
	tree.Insert(12, 12)
	tree.Insert(7, 7)
	tree.Insert(16, 16)
	tree.Insert(4, 4)
	tree.Insert(6, 6)
	tree.Insert(8, 8)
	tree.Insert(9, 9)
	tree.Insert(10, 10)

	tree.Visualize("testInsert")
}

func TestRedBlackTreeDelete(t *testing.T) {
	// generate rbt
	tree := rbt.NewRBTree()

	tree.Insert(12, 12)
	tree.Insert(7, 7)
	tree.Insert(16, 16)
	tree.Insert(4, 4)
	tree.Insert(6, 6)
	tree.Insert(8, 8)
	tree.Insert(9, 9)
	tree.Insert(10, 10)

	t.Run("normal delete", func(t *testing.T) {
		tree.Delete(6)
		tree.Visualize("testDelete")
	})
}
