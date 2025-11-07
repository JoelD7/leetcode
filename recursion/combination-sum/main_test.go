package combination_sum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	c := require.New(t)

	t.Run("candidates = [2,3,6,7], target = 7", func(t *testing.T) {
		c.ElementsMatch([][]int{{2, 2, 3}, {7}}, combinationSum([]int{2, 3, 6, 7}, 7))
	})

	t.Run("candidates = [2,3,5], target = 8", func(t *testing.T) {
		c.ElementsMatch([][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}, combinationSum([]int{2, 3, 5}, 8))
	})

	t.Run("candidates = [2], target = 1", func(t *testing.T) {
		c.Empty(combinationSum([]int{2}, 1))
	})
}
