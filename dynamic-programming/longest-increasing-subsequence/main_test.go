package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLIS(t *testing.T) {
	t.Run("Standard mixed array", func(t *testing.T) {
		nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
		assert.Equal(t, 4, lengthOfLIS(nums))
	})

	t.Run("Array with zeros and repeating numbers", func(t *testing.T) {
		nums := []int{0, 1, 0, 3, 2, 3}
		assert.Equal(t, 4, lengthOfLIS(nums))
	})

	t.Run("Array with all identical elements", func(t *testing.T) {
		nums := []int{7, 7, 7, 7, 7, 7, 7}
		assert.Equal(t, 1, lengthOfLIS(nums))
	})

	t.Run("Strictly increasing array", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		assert.Equal(t, 5, lengthOfLIS(nums))
	})

	t.Run("Strictly decreasing array", func(t *testing.T) {
		nums := []int{5, 4, 3, 2, 1}
		assert.Equal(t, 1, lengthOfLIS(nums))
	})

	t.Run("Array with negative numbers", func(t *testing.T) {
		nums := []int{-3, -2, -1, -5, -4, 0}
		assert.Equal(t, 4, lengthOfLIS(nums))
	})

	t.Run("Single element array", func(t *testing.T) {
		nums := []int{42}
		assert.Equal(t, 1, lengthOfLIS(nums))
	})
}
