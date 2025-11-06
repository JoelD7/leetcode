package subsets

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	comb := make([]int, 0)

	backtrack(0, nums, &comb, &res)

	return res
}

func backtrack(start int, nums []int, comb *[]int, res *[][]int) {
	combCopy := make([]int, len(*comb))
	copy(combCopy, *comb)
	*res = append(*res, combCopy)

	for i := start; i < len(nums); i++ {
		*comb = append(*comb, nums[i])
		backtrack(i+1, nums, comb, res)
		*comb = combCopy[:len(*comb)-1]
	}
}
