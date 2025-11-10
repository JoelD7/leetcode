package combination_sum_2

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	comb := make([]int, 0)

	sort.Ints(candidates)

	res = backtrack(0, target, 0, candidates, comb, res)

	return res
}

func backtrack(start, target, sum int, candidates, comb []int, res [][]int) [][]int {
	if sum == target {
		c := make([]int, len(comb))
		copy(c, comb)
		res = append(res, c)
		return res
	}

	if sum > target {
		return res
	}

	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}

		comb = append(comb, candidates[i])
		sum += candidates[i]

		res = backtrack(i+1, target, sum, candidates, comb, res)

		comb = comb[:len(comb)-1]
		sum -= candidates[i]
	}

	return res
}
