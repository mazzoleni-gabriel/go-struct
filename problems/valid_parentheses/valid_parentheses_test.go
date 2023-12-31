package valid_parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidParentheses(t *testing.T) {

	t.Run("()", func(t *testing.T) {
		assert.True(t, IsValid("()"))
	})

	t.Run("()[]{}", func(t *testing.T) {
		assert.True(t, IsValid("()[]{}"))
	})

	t.Run("(]", func(t *testing.T) {
		assert.False(t, IsValid("(]"))
	})

	t.Run("({)}", func(t *testing.T) {
		assert.False(t, IsValid("({)}"))
	})
}
