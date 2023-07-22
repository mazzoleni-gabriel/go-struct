package quicksort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {

	t.Run("sort random", func(t *testing.T) {
		sorter := New([]int{3, 7, 1, 2, 9, 5, 4})
		sorter.Sort()
		assert.Equal(t, []int{1, 2, 3, 4, 5, 7, 9}, sorter.Get())
	})

	t.Run("sort already sorted", func(t *testing.T) {
		sorter := New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
		sorter.Sort()
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, sorter.Get())
	})

	t.Run("sort desc sorted", func(t *testing.T) {
		sorter := New([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
		sorter.Sort()
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, sorter.Get())
	})

	t.Run("sort even random", func(t *testing.T) {
		sorter := New([]int{3, 7, 1, 9, 5, 4})
		sorter.Sort()
		assert.Equal(t, []int{1, 3, 4, 5, 7, 9}, sorter.Get())
	})

}
