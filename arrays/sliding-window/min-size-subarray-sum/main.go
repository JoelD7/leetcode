package min_size_subarray_sum

import (
	"math"
)

// target = 11, nums = [1,2,3,4,5]
//
//	prefixSum =       [1,3,6,10,15]
func minSubArrayLen(target int, nums []int) int {
	prefix := prefixSum(nums)
	min := func(a, b int) int {
		if a < b {
			return a
		}

		return b
	}
	var i, j, sum int
	minLength := math.MaxInt32

	if prefix[len(prefix)-1] < target {
		return 0
	}

	for ; j < len(prefix); j++ {
		//All the numbers after this condition is satisfied for the first time are >= target
		if prefix[j] >= target && i == 0 {
			minLength = min(minLength, j+1)
			i = j + 1
			break
		}
	}

	for ; i < len(nums); i++ {
		sum += nums[i]
		if sum >= target {
			minLength = min(minLength, i-j)
			i++
		}
	}

	if minLength == math.MaxInt32 {
		return 0
	}

	return minLength
}

func prefixSum(nums []int) []int {
	prefix := make([]int, len(nums))
	prefix[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		prefix[i] = nums[i] + prefix[i-1]
	}

	return prefix
}
