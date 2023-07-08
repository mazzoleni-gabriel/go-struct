package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {

	t.Run("pop", func(t *testing.T) {
		queue := Queue{}
		queue.Push(10)
		assert.Equal(t, 10, queue.Pop())
		assert.True(t, queue.IsEmpty())
	})

	t.Run("top", func(t *testing.T) {
		queue := Queue{}
		queue.Push(10)
		assert.Equal(t, 10, queue.Top())
		assert.False(t, queue.IsEmpty())
	})

}
