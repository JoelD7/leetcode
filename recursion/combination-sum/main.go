package combination_sum

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	comb := make([]int, 0)
	var sum int

	res = backtrack(0, sum, target, candidates, comb, res)

	return res
}

func backtrack(start, sum, target int, candidates, comb []int, res [][]int) [][]int {
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
		comb = append(comb, candidates[i])
		sum += candidates[i]

		res = backtrack(i, sum, target, candidates, comb, res)

		comb = comb[:len(comb)-1]
		sum -= candidates[i]
	}

	return res
}
