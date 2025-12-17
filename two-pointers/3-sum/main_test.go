package three_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSum(t *testing.T) {

	t.Run("nums = [-1,0,1,2,-1,-4]", func(t *testing.T) {
		assert.ElementsMatch(t, threeSum([]int{-1, 0, 1, 2, -1, -4}), [][]int{{-1, -1, 2}, {-1, 0, 1}})
	})

	t.Run("nums = [0,1,1]", func(t *testing.T) {
		assert.Empty(t, threeSum([]int{0, 1, 1}))
	})

	t.Run("nums = [0,0,0]", func(t *testing.T) {
		assert.ElementsMatch(t, threeSum([]int{0, 0, 0}), [][]int{{0, 0, 0}})
	})

	t.Run("[0,0,0,0]", func(t *testing.T) {
		assert.ElementsMatch(t, threeSum([]int{0, 0, 0, 0}), [][]int{{0, 0, 0}})
	})

	t.Run("[1,2,-2,-1]", func(t *testing.T) {
		assert.ElementsMatch(t, threeSum([]int{1, 2, -2, -1}), [][]int{})
	})
}
