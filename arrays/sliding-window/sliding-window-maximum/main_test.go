package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSlidingWindow(t *testing.T) {
	//Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
	//	Output: [3,3,5,5,6,7]

	t.Run("nums = [1,3,-1,-3,5,3,6,7], k = 3", func(t *testing.T) {
		assert.Equal(t, []int{3, 3, 5, 5, 6, 7}, maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	})

	t.Run("nums = 1, k = 1", func(t *testing.T) {
		assert.Equal(t, []int{1}, maxSlidingWindow([]int{1}, 1))
	})
}
