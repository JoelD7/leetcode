package combination_sum

func combinationSum(candidates []int, target int) [][]int {
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
			comb = append(comb, candidates[i])
			backtrack(i, sum+candidates[i])
			comb = comb[:len(comb)-1]
		}
	}

	backtrack(0, 0)
	return res
}
