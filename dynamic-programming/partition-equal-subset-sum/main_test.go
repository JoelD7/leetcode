package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanPartition(t *testing.T) {
	t.Run("nums = [1, 5, 11, 5] (can be partitioned into [1, 5, 5] and [11])", func(t *testing.T) {
		nums := []int{1, 5, 11, 5}
		assert.Equal(t, true, canPartition(nums))
	})

	t.Run("nums = [1, 2, 3, 5] (cannot be partitioned)", func(t *testing.T) {
		nums := []int{1, 2, 3, 5}
		assert.Equal(t, false, canPartition(nums))
	})

	t.Run("odd sum array cannot be partitioned", func(t *testing.T) {
		nums := []int{1, 2, 4}
		assert.Equal(t, false, canPartition(nums))
	})

	t.Run("array with two identical elements", func(t *testing.T) {
		nums := []int{2, 2}
		assert.Equal(t, true, canPartition(nums))
	})

	t.Run("array with a single element", func(t *testing.T) {
		nums := []int{10}
		assert.Equal(t, false, canPartition(nums))
	})

	t.Run("larger array with valid partition", func(t *testing.T) {
		nums := []int{14, 9, 8, 4, 3, 2}
		assert.Equal(t, true, canPartition(nums))
	})

	t.Run("zeroes in the array", func(t *testing.T) {
		nums := []int{2, 0, 2}
		assert.Equal(t, true, canPartition(nums))
	})

	t.Run("[3,3,6,8,16,16,16,18,20]", func(t *testing.T) {
		nums := []int{3, 3, 6, 8, 16, 16, 16, 18, 20}
		assert.Equal(t, true, canPartition(nums))
	})

	t.Run("[23,13,11,7,6,5,5]", func(t *testing.T) {
		nums := []int{23, 13, 11, 7, 6, 5, 5}
		assert.Equal(t, true, canPartition(nums))
	})
}
