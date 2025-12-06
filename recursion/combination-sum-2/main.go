package combination_sum_2

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	comb := make([]int, 0)

	var backtrack func(start, sum int)

	backtrack = func(start, sum int) {
		if sum > target {
			return
		}

		if sum == target {
			res = append(res, append([]int(nil), comb...))
			return
		}

		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i-1] == candidates[i] {
				continue
			}
			comb = append(comb, candidates[i])
			backtrack(i+1, sum+candidates[i])
			comb = comb[:len(comb)-1]
		}
	}

	backtrack(0, 0)
	return res
}
