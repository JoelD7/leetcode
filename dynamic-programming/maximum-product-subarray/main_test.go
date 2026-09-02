package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProduct(t *testing.T) {
	t.Run("Example 1 from Leetcode", func(t *testing.T) {
		nums := []int{2, 3, -2, 4}
		assert.Equal(t, 6, maxProduct(nums))
	})

	t.Run("Custom examle", func(t *testing.T) {
		nums := []int{-2, -8, -2, 4}
		assert.Equal(t, 64, maxProduct(nums))
	})

	t.Run("Example 2 from Leetcode", func(t *testing.T) {
		nums := []int{-2, 0, -1}
		assert.Equal(t, 0, maxProduct(nums))
	})

	t.Run("Single negative element", func(t *testing.T) {
		nums := []int{-2}
		assert.Equal(t, -2, maxProduct(nums))
	})

	t.Run("Two negative numbers resulting in a positive", func(t *testing.T) {
		nums := []int{-2, -3}
		assert.Equal(t, 6, maxProduct(nums))
	})

	t.Run("Alternating positive and negative numbers", func(t *testing.T) {
		nums := []int{-2, 3, -4}
		assert.Equal(t, 24, maxProduct(nums))
	})

	t.Run("All negative numbers", func(t *testing.T) {
		nums := []int{-2, -3, -4}
		// The max product would be from -3 * -4 = 12
		assert.Equal(t, 12, maxProduct(nums))
	})

	t.Run("Array with zero in the middle", func(t *testing.T) {
		nums := []int{1, 2, 0, 3, 4}
		assert.Equal(t, 12, maxProduct(nums))
	})

	t.Run("Large continuous product", func(t *testing.T) {
		nums := []int{2, 3, 2, 4}
		assert.Equal(t, 48, maxProduct(nums))
	})

	t.Run("0,2", func(t *testing.T) {
		nums := []int{0, 2}
		assert.Equal(t, 2, maxProduct(nums))
	})

	t.Run("[-3,0,1,-2]", func(t *testing.T) {
		nums := []int{-3, 0, 1, -2}
		assert.Equal(t, 1, maxProduct(nums))
	})
}
