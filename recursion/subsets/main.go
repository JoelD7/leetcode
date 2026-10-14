package main

func subsets(nums []int) [][]int {
	comb := make([]int, 0)
	res := make([][]int, 0)

	var recurse func(start int)
	recurse = func(start int) {
		combCopy := make([]int, len(comb))
		copy(combCopy, comb)

		res = append(res, combCopy)

		if start == len(nums) {
			return
		}

		for i := start; i < len(nums); i++ {
			comb = append(comb, nums[i])
			recurse(i + 1)
			comb = comb[:len(comb)-1]
		}
	}

	recurse(0)
	return res
}
