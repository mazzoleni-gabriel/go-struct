package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {

	t.Run("pop", func(t *testing.T) {
		stack := Stack{}
		stack.Push(10)
		stack.Push(11)
		stack.Push(12)
		assert.Equal(t, 12, stack.Pop())
		assert.Equal(t, 11, stack.Pop())
		assert.Equal(t, 10, stack.Pop())
		assert.True(t, stack.IsEmpty())
	})

	t.Run("top", func(t *testing.T) {
		stack := Stack{}
		stack.Push(10)
		stack.Push(11)
		assert.Equal(t, 11, stack.Top())
		assert.Equal(t, 11, stack.Top())
		assert.False(t, stack.IsEmpty())
	})

}
