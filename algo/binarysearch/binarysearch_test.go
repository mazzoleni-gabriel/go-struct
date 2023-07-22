package binarysearch

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {

	t.Run("find first element", func(t *testing.T) {
		searcher := New([]int{1, 2, 3, 4, 5, 6, 7})
		assert.True(t, searcher.Search(1))
	})

	t.Run("find last element", func(t *testing.T) {
		searcher := New([]int{1, 2, 3, 4, 5, 6, 7})
		assert.True(t, searcher.Search(7))
	})

	t.Run("find mid element", func(t *testing.T) {
		searcher := New([]int{1, 2, 3, 4, 5, 6, 7})
		assert.True(t, searcher.Search(4))
	})

	t.Run("find random element", func(t *testing.T) {
		searcher := New([]int{1, 2, 3, 4, 5, 6, 7})
		assert.True(t, searcher.Search(5))
	})

	t.Run("find odd list", func(t *testing.T) {
		searcher := New([]int{1, 2, 3, 5, 6, 7})
		assert.True(t, searcher.Search(5))
	})

	t.Run("dont find bigger element", func(t *testing.T) {
		searcher := New([]int{1, 2, 3, 4, 5, 6, 7})
		assert.False(t, searcher.Search(10))
	})

	t.Run("dont find smaller element", func(t *testing.T) {
		searcher := New([]int{1, 2, 3, 4, 5, 6, 7})
		assert.False(t, searcher.Search(-10))
	})

	t.Run("dont find mid element", func(t *testing.T) {
		searcher := New([]int{1, 3, 4, 5, 6, 7})
		assert.False(t, searcher.Search(2))
	})

}
