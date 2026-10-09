package container_most_water

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	t.Run("Leetcode Example 1", func(t *testing.T) {
		assert.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	})

	t.Run("Leetcode Example 2", func(t *testing.T) {
		assert.Equal(t, 1, maxArea([]int{1, 1}))
	})

	t.Run("Decreasing heights", func(t *testing.T) {
		assert.Equal(t, 6, maxArea([]int{5, 4, 3, 2, 1}))
	})

	t.Run("Increasing heights", func(t *testing.T) {
		assert.Equal(t, 6, maxArea([]int{1, 2, 3, 4, 5}))
	})

	t.Run("All elements are the same", func(t *testing.T) {
		assert.Equal(t, 24, maxArea([]int{8, 8, 8, 8}))
	})
}
