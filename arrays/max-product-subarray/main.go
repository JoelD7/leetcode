package max_product_subarray

import (
	"math"
)

func maxProduct(nums []int) int {
	maxDot := math.MinInt32
	curDot := 1
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 0; i < len(nums); i++ {
		curDot *= nums[i]
		maxDot = max(curDot, maxDot)
	}

	curDot = 1

	for i := len(nums) - 1; i >= 0; i-- {
		curDot *= nums[i]
		maxDot = max(curDot, maxDot)
	}

	return maxDot
}
