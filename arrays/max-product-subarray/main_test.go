package max_product_subarray

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	c := require.New(t)

	t.Run("2,3,-2,4", func(t *testing.T) {
		c.Equal(6, maxProduct([]int{2, 3, -2, 4}))
	})

	t.Run("-2,0,-1", func(t *testing.T) {
		c.Equal(0, maxProduct([]int{-2, 0, -1}))
	})

	t.Run("0,2", func(t *testing.T) {
		c.Equal(2, maxProduct([]int{0, 2}))
	})

	t.Run("-2,3,-4", func(t *testing.T) {
		c.Equal(24, maxProduct([]int{-2, 3, -4}))
	})

	t.Run("2,-5,-2,-4,3", func(t *testing.T) {
		c.Equal(24, maxProduct([]int{2, -5, -2, -4, 3}))
	})

	t.Run("-3,0,1,-2", func(t *testing.T) {
		c.Equal(1, maxProduct([]int{-3, 0, 1, -2}))
	})
}
