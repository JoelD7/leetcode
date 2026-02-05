package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSlidingWindow(t *testing.T) {
	t.Run("nums = [1,3,-1,-3,5,3,6,7], k = 3", func(t *testing.T) {
		assert.Equal(t, []int{3, 3, 5, 5, 6, 7}, maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	})

	t.Run("nums = [1,3,1,2,0,5], k = 3", func(t *testing.T) {
		assert.Equal(t, []int{3, 3, 2, 5}, maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))
	})

	t.Run("nums = [9,10,9,-7,-4,-8,2,-6], k = 5", func(t *testing.T) {
		assert.Equal(t, []int{10, 10, 9, 2}, maxSlidingWindow([]int{9, 10, 9, -7, -4, -8, 2, -6}, 5))
	})

	t.Run("nums = 1, k = 1", func(t *testing.T) {
		assert.Equal(t, []int{1}, maxSlidingWindow([]int{1}, 1))
	})

	t.Run("nums = 1,-1, k = 1", func(t *testing.T) {
		assert.Equal(t, []int{1, -1}, maxSlidingWindow([]int{1, -1}, 1))
	})
}
