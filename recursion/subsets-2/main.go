package subsets_2

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{{}}
	comb := make([]int, 0)
	return backtrack(0, nums, comb, res)
}

// [1,2,2]
func backtrack(start int, nums, comb []int, res [][]int) [][]int {
	for i := start; i < len(nums); i++ {
		comb = append(comb, nums[i])
		c := make([]int, len(comb))
		copy(c, comb)
		res = append(res, c)

		res = backtrack(i+1, nums, comb, res)
		comb = []int{}
	}

	return res
}
