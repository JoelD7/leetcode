package min_size_subarray_sum

import (
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}

		return b
	}
	var sum, i int
	minLength := math.MaxInt32

	for j := 0; j < len(nums); j++ {
		sum += nums[j]

		for sum >= target {
			minLength = min(minLength, j-i+1)
			sum -= nums[i]
			i++
		}
	}

	if minLength == math.MaxInt32 {
		return 0
	}

	return minLength
}
