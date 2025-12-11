package subsets_2

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	comb := make([]int, 0)
	sort.Ints(nums)

	var backtrack = func(i int) {}

	backtrack = func(i int) {
		if i == len(nums) {
			res = append(res, append([]int{}, comb...))
			return
		}

		comb = append(comb, nums[i])
		backtrack(i + 1)

		comb = comb[:len(comb)-1]
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
		backtrack(i + 1)
	}

	backtrack(0)
	return res
}
