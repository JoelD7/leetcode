package main

// input: -2,1,-3,4,-1,2,1,-5,4
// output: 6 (4,-1,2,1)
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	max := -10000
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum = nums[i]
		if sum > max {
			max = sum
		}

		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum > max {
				max = sum
			}
		}
	}

	return max
}
