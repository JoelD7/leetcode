package valid_parentheses

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsValid(t *testing.T) {
	c := require.New(t)

	t.Run("]", func(t *testing.T) {
		c.False(isValid("]"))
	})

	t.Run(")(){}", func(t *testing.T) {
		c.False(isValid(")(){}"))
	})

	t.Run("([}}])", func(t *testing.T) {
		c.False(isValid("([}}])"))
	})

	c.False(isValid("(("))
	c.True(isValid("{[]}"))
	c.True(isValid("()"))
	c.True(isValid("()[]{}"))
}
