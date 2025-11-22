package subsets_2

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	comb := make([]int, 0)
	sort.Ints(nums)

	return backtrack(0, nums, comb, res)
}

func backtrack(i int, nums, comb []int, res [][]int) [][]int {
	if i == len(nums) {
		c := make([]int, len(comb))
		copy(c, comb)
		res = append(res, c)
		return res
	}

	comb = append(comb, nums[i])
	res = backtrack(i+1, nums, comb, res)

	comb = comb[:len(comb)-1]

	for i+1 < len(nums) && nums[i] == nums[i+1] {
		i++
	}

	return backtrack(i+1, nums, comb, res)
}
