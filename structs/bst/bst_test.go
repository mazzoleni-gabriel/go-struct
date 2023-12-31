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
		//          10
		//         /  \
		//        5    15
		//       / \   / \
		//      3   7 13  17

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

func TestBST_Search(t *testing.T) {

	t.Run("search existing", func(t *testing.T) {
		//          10
		//         /  \
		//        5    15
		//       / \   / \
		//      3   7 13  17
		bst := BST{}

		bst.Insert(10)
		bst.Insert(5)
		bst.Insert(3)
		bst.Insert(7)
		bst.Insert(15)
		bst.Insert(17)
		bst.Insert(13)

		assert.True(t, bst.Search(13))
	})

	t.Run("search non existing", func(t *testing.T) {
		//          10
		//         /  \
		//        5    15
		//       / \   / \
		//      3   7 13  17
		bst := BST{}

		bst.Insert(10)
		bst.Insert(5)
		bst.Insert(3)
		bst.Insert(7)
		bst.Insert(15)
		bst.Insert(17)
		bst.Insert(13)

		assert.False(t, bst.Search(14))
	})

}

func TestBST_Max(t *testing.T) {

	t.Run("max found", func(t *testing.T) {
		//          10
		//         /  \
		//        5    15
		//       / \   / \
		//      3   7 13  17
		bst := BST{}

		bst.Insert(10)
		bst.Insert(5)
		bst.Insert(3)
		bst.Insert(7)
		bst.Insert(15)
		bst.Insert(17)
		bst.Insert(13)

		assert.Equal(t, 17, bst.Max())
	})

	t.Run("max empty BST", func(t *testing.T) {
		bst := BST{}

		assert.Equal(t, 0, bst.Max())
	})

}

func TestBST_Min(t *testing.T) {

	t.Run("min found", func(t *testing.T) {
		//          10
		//         /  \
		//        5    15
		//       / \   / \
		//      3   7 13  17
		bst := BST{}

		bst.Insert(10)
		bst.Insert(5)
		bst.Insert(3)
		bst.Insert(7)
		bst.Insert(15)
		bst.Insert(17)
		bst.Insert(13)

		assert.Equal(t, 3, bst.Min())
	})

	t.Run("min empty BST", func(t *testing.T) {
		bst := BST{}

		assert.Equal(t, 0, bst.Min())
	})

}

func TestBST_BFS(t *testing.T) {

	t.Run("bfs", func(t *testing.T) {
		//          10
		//         /  \
		//        5    15
		//       / \   / \
		//      3   7 13  17
		bst := BST{}

		bst.Insert(10)
		bst.Insert(5)
		bst.Insert(3)
		bst.Insert(7)
		bst.Insert(15)
		bst.Insert(17)
		bst.Insert(13)

		assert.Equal(t, []int{10, 5, 15, 3, 7, 13, 17}, bst.BFS())
	})
}

func TestBST_DFS(t *testing.T) {

	t.Run("dfs", func(t *testing.T) {
		//          10
		//         /  \
		//        5    15
		//       / \   / \
		//      3   7 13  17
		bst := BST{}

		bst.Insert(10)
		bst.Insert(5)
		bst.Insert(3)
		bst.Insert(7)
		bst.Insert(15)
		bst.Insert(17)
		bst.Insert(13)

		assert.Equal(t, []int{10, 5, 3, 7, 15, 13, 17}, bst.DFS())
	})
}

func TestBST_Height(t *testing.T) {

	t.Run("height 3", func(t *testing.T) {
		//          10
		//         /  \
		//        5    15
		//       / \   / \
		//      3   7 13  17
		bst := BST{}

		bst.Insert(10)
		bst.Insert(5)
		bst.Insert(3)
		bst.Insert(7)
		bst.Insert(15)
		bst.Insert(17)
		bst.Insert(13)

		assert.Equal(t, 3, bst.Height())
	})

	t.Run("height 4 left", func(t *testing.T) {
		//          10
		//         /
		//        5
		//       /
		//      3
		//     /
		//    2
		bst := BST{}

		bst.Insert(10)
		bst.Insert(5)
		bst.Insert(3)
		bst.Insert(2)

		assert.Equal(t, 4, bst.Height())
	})

	t.Run("height 4 right", func(t *testing.T) {
		//          10
		//            \
		//             15
		//               \
		//                17
		//                 \
		//                  20
		bst := BST{}

		bst.Insert(10)
		bst.Insert(15)
		bst.Insert(17)
		bst.Insert(20)

		assert.Equal(t, 4, bst.Height())
	})

	t.Run("height 1", func(t *testing.T) {
		//          10
		bst := BST{}

		bst.Insert(10)

		assert.Equal(t, 1, bst.Height())
	})

	t.Run("height 0", func(t *testing.T) {
		bst := BST{}

		assert.Equal(t, 0, bst.Height())
	})
}
