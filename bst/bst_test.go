package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBST_Insert(t *testing.T) {

	t.Run("insert one node", func(t *testing.T) {
		bst := BST{}
		bst.Insert(10)
		assert.Equal(t, 10, bst.root.data)
	})

	t.Run("insert right node", func(t *testing.T) {
		bst := BST{}

		bst.Insert(10)
		bst.Insert(15)

		assert.Equal(t, 10, bst.root.data)
		assert.Equal(t, 15, bst.root.right.data)
	})

	t.Run("insert left node", func(t *testing.T) {
		bst := BST{}

		bst.Insert(10)
		bst.Insert(5)

		assert.Equal(t, 10, bst.root.data)
		assert.Equal(t, 5, bst.root.left.data)
	})

	t.Run("insert tree", func(t *testing.T) {
		bst := BST{}

		bst.Insert(10)
		bst.Insert(5)
		bst.Insert(3)
		bst.Insert(7)
		bst.Insert(15)
		bst.Insert(17)
		bst.Insert(13)

		assert.Equal(t, 10, bst.root.data)
		assert.Equal(t, 5, bst.root.left.data)
		assert.Equal(t, 3, bst.root.left.left.data)
		assert.Equal(t, 7, bst.root.left.right.data)
		assert.Equal(t, 15, bst.root.right.data)
		assert.Equal(t, 13, bst.root.right.left.data)
		assert.Equal(t, 17, bst.root.right.right.data)
	})

}
