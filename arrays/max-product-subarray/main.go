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
		curDot = nums[i]
		maxDot = max(maxDot, curDot)
		if nums[i] == 0 {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			curDot *= nums[j]
			maxDot = max(maxDot, curDot)
		}
	}

	return maxDot
}
