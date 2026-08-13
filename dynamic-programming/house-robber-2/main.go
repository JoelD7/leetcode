package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var robHelper func(start, end int, nums []int) int
	robHelper = func(start, end int, nums []int) int {
		if len(nums[start:end+1]) == 1 {
			return nums[start]
		}

		memo := make([]int, end+1)

		memo[start] = nums[start]
		memo[start+1] = max(nums[start+1], memo[start])
		for i := start + 2; i <= end; i++ {
			memo[i] = max(memo[i-2]+nums[i], memo[i-1])
		}

		return memo[end]
	}

	return max(robHelper(0, len(nums)-2, nums), robHelper(1, len(nums)-1, nums))
}
