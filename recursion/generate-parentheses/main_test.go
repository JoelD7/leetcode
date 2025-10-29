package generate_parentheses

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	c := require.New(t)

	t.Run("n = 3", func(t *testing.T) {
		c.Equal([]string{"((()))", "(()())", "(())()", "()(())", "()()()"}, generateParenthesis(3))
	})

	t.Run("n = 1", func(t *testing.T) {
		c.Equal([]string{"()"}, generateParenthesis(1))
	})
}
