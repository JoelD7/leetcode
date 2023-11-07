package arrays

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindSubarrayZeroSum(t *testing.T) {
	c := require.New(t)

	c.True(findSubarrayZeroSum([]int{4, 2, -3, 1, 6}))
	c.True(findSubarrayZeroSum([]int{5, 1, 2, -1, -1, 6}))
	c.True(findSubarrayZeroSum([]int{8, 9, 3, 1, -6, 6}))
	c.True(findSubarrayZeroSum([]int{4, 2, 0, 1, 6}))
	c.True(findSubarrayZeroSum([]int{4, 2, 1, 6, 6, 3, 4, 5, 6, 7, 8, -10, -5, -11, 9}))
}
