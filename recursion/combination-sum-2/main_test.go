package combination_sum_2

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	c := require.New(t)

	t.Run("candidates = [10,1,2,7,6,1,5], target = 8", func(t *testing.T) {
		c.ElementsMatch([][]int{{1, 1, 6}, {2, 1, 5}, {7, 1}, {2, 6}}, combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	})

	t.Run("candidates = [2,5,2,1,2], target = 5", func(t *testing.T) {
		c.ElementsMatch([][]int{{2, 1, 2}, {5}}, combinationSum2([]int{2, 5, 2, 1, 2}, 5))
	})
}
