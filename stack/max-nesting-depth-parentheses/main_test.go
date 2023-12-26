package max_nesting_depth_parentheses

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	c := require.New(t)

	c.Equal(3, maxDepth("(1+(2*3)+((8)/4))+1"))
	c.Equal(3, maxDepth("(1)+((2))+(((3)))"))
}
