package largest_rectangle_in_histogram

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	t.Run("heights = [2,1,5,6,2,3]", func(t *testing.T) {
		assert.Equal(t, 10, largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	})

	t.Run("heights = [2,4]", func(t *testing.T) {
		assert.Equal(t, 4, largestRectangleArea([]int{2, 4}))
	})

	t.Run("heights = [4,2]", func(t *testing.T) {
		assert.Equal(t, 4, largestRectangleArea([]int{4, 2}))
	})

	t.Run("heights = [0]", func(t *testing.T) {
		assert.Equal(t, 0, largestRectangleArea([]int{0}))
	})

	t.Run("heights = [0,9]", func(t *testing.T) {
		assert.Equal(t, 0, largestRectangleArea([]int{0}))
	})
}
