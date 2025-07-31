package longest_harmonious_subsequence

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindLHS(t *testing.T) {
	c := require.New(t)

	t.Run("[2,3,3,2,-2,0,3,3,3,1,-2,1,0,1,2,0,2,2,2,1,2,2,1,1,2,2]", func(t *testing.T) {
		c.Equal(16, findLHS([]int{2, 3, 3, 2, -2, 0, 3, 3, 3, 1, -2, 1, 0, 1, 2, 0, 2, 2, 2, 1, 2, 2, 1, 1, 2, 2}))
	})

	t.Run("[1,2,2,1]", func(t *testing.T) {
		c.Equal(4, findLHS([]int{1, 2, 2, 1}))
	})

	t.Run("[1,3,5,7,9,11,13,15,17]", func(t *testing.T) {
		c.Equal(0, findLHS([]int{1, 3, 5, 7, 9, 11, 13, 15, 17}))
	})

	t.Run("[1,1,1,1]", func(t *testing.T) {
		c.Equal(0, findLHS([]int{1, 1, 1, 1}))
	})

	t.Run("[1,3,2,2,5,2,3,7]", func(t *testing.T) {
		c.Equal(5, findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
	})

	t.Run("[1,2,3,4]", func(t *testing.T) {
		c.Equal(2, findLHS([]int{1, 2, 3, 4}))
	})

}
