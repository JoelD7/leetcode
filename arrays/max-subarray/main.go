package main

func maxSubArray(nums []int) int {
	maxSum := -10000 //Min value according to the problem description
	curMaxSum := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < len(nums); i++ {
		if nums[i]+curMaxSum >= nums[i] {
			maxSum = max(maxSum, nums[i]+curMaxSum)
			curMaxSum += nums[i]
			continue
		}

		curMaxSum = nums[i]
		maxSum = max(maxSum, nums[i])
	}

	return maxSum
}
