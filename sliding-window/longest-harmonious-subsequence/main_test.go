package longest_harmonious_subsequence

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindLHS(t *testing.T) {
	c := require.New(t)

	t.Run("[1,3,2,2,5,2,3,7]", func(t *testing.T) {
		c.Equal(5, findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
	})

	t.Run("[1,2,3,4]", func(t *testing.T) {
		c.Equal(2, findLHS([]int{1, 2, 3, 4}))
	})

	t.Run("[1,1,1,1]", func(t *testing.T) {
		c.Equal(0, findLHS([]int{1, 1, 1, 1}))
	})
}
