package combination_sum_2

import (
	"sort"
	"strconv"
)

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	comb := make([]int, 0)
	combsMade := make(map[string]struct{})
	combKey := ""

	sort.Ints(candidates)

	res = backtrack(0, target, 0, candidates, comb, combKey, combsMade, res)

	return res
}

func backtrack(start, target, sum int, candidates, comb []int, combKey string, combsMade map[string]struct{}, res [][]int) [][]int {
	if sum == target {
		if _, ok := combsMade[combKey]; ok {
			return res
		}

		c := make([]int, len(comb))
		copy(c, comb)
		combsMade[combKey] = struct{}{}
		res = append(res, c)
		return res
	}

	if sum > target {
		return res
	}

	for i := start; i < len(candidates); i++ {
		comb = append(comb, candidates[i])
		sum += candidates[i]
		combKey += strconv.Itoa(candidates[i])

		res = backtrack(i+1, target, sum, candidates, comb, combKey, combsMade, res)

		comb = comb[:len(comb)-1]
		sum -= candidates[i]
		combKey = backtrackCombKey(combKey, candidates[i])
	}

	return res
}

func backtrackCombKey(combKey string, candidate int) string {
	candidateStr := strconv.Itoa(candidate)
	return combKey[:len(combKey)-len(candidateStr)]
}
