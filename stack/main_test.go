package stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsValid(t *testing.T) {
	c := require.New(t)

	c.False(isValid("(("))
	c.True(isValid("()"))
	c.True(isValid("()[]{}"))
}
