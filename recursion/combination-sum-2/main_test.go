package combination_sum_2

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	c := require.New(t)

	t.Run("candidates = [14,6,25,9,30,20,33,34,28,30,16,12,31,9,9,12,34,16,25,32,8,7,30,12,33,20,21,29,24,17,27,34,11,17,30,6,32,21,27,17,16,8,24,12,12,28,11,33,10,32,22,13,34,18,12] target = 27", func(t *testing.T) {
		candidates := []int{14, 6, 25, 9, 30, 20, 33, 34, 28, 30, 16, 12, 31, 9, 9, 12, 34, 16, 25, 32, 8, 7, 30, 12, 33, 20, 21, 29, 24, 17, 27, 34, 11, 17, 30, 6, 32, 21, 27, 17, 16, 8, 24, 12, 12, 28, 11, 33, 10, 32, 22, 13, 34, 18, 12}
		c.ElementsMatch([][]int{
			{6, 6, 7, 8},
			{6, 7, 14},
			{6, 8, 13},
			{6, 9, 12},
			{6, 10, 11},
			{6, 21},
			{7, 8, 12},
			{7, 9, 11},
			{7, 20},
			{8, 8, 11},
			{8, 9, 10},
			{9, 9, 9},
			{9, 18},
			{10, 17},
			{11, 16},
			{13, 14},
			{27},
		}, combinationSum2(candidates, 27))
	})

	t.Run("candidates = [10,1,2,7,6,1,5], target = 8", func(t *testing.T) {
		c.ElementsMatch([][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}, combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	})

	t.Run("candidates = [2,5,2,1,2], target = 5", func(t *testing.T) {
		c.ElementsMatch([][]int{{1, 2, 2}, {5}}, combinationSum2([]int{2, 5, 2, 1, 2}, 5))
	})
}
