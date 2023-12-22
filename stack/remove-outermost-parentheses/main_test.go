package remove_outermost_parentheses

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveOutermostParentheses(t *testing.T) {
	c := require.New(t)

	c.Equal("()()()", removeOuterParentheses("(()())(())"))
	c.Equal("()()()()(())", removeOuterParentheses("(()())(())(()(()))"))
	c.Equal("", removeOuterParentheses("()()"))
}
