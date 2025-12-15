package longest_consecutive_sequence

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	c := require.New(t)

	c.Equal(3, longestConsecutive([]int{1, 2, 0, 1}))
	c.Equal(4, longestConsecutive([]int{9, 1, -3, 2, 4, 8, 3, -1, 6, -2, -4, 7}))
	c.Equal(9, longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	c.Equal(4, longestConsecutive([]int{100, 4, 200, 1, 3, 2}))

}
