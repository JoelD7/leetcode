package subsets

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	comb := make([]int, 0)

	var backtrack func(start int)

	backtrack = func(start int) {
		res = append(res, append([]int(nil), comb...))

		for i := start; i < len(nums); i++ {
			comb = append(comb, nums[i])
			backtrack(i + 1)
			comb = comb[:len(comb)-1]
		}
	}

	backtrack(0)
	return res
}
