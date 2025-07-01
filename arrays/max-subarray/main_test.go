package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	c := require.New(t)

	t.Run("-2,1,-3,4,-1,2,1,-5,4", func(t *testing.T) {
		c.Equal(6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	})

	t.Run("1", func(t *testing.T) {
		c.Equal(1, maxSubArray([]int{1}))
	})

	t.Run("1,2", func(t *testing.T) {
		c.Equal(3, maxSubArray([]int{1, 2}))
	})

	t.Run("5,4,-1,7,8", func(t *testing.T) {
		c.Equal(23, maxSubArray([]int{5, 4, -1, 7, 8}))
	})
}
