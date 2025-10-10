package min_size_subarray_sum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMinSubarrayLen(t *testing.T) {
	c := require.New(t)

	t.Run("target = 7, nums = [2,3,1,2,4,3]", func(t *testing.T) {
		c.Equal(2, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	})

	t.Run("target = 4, nums = [1,4,4]", func(t *testing.T) {
		c.Equal(1, minSubArrayLen(4, []int{1, 4, 4}))
	})

	t.Run("target = 11, nums = [1,1,1,1,1,1,1,1]", func(t *testing.T) {
		c.Equal(0, minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
	})

	t.Run("target = 11, nums = [1,2,3,4,5]", func(t *testing.T) {
		c.Equal(3, minSubArrayLen(11, []int{1, 2, 3, 4, 5}))
	})
}
