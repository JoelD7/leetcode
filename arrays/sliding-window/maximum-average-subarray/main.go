package maximum_average_subarray

import (
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	i := 0
	j := 0
	var total float64
	maxAvg := float64(math.MinInt64)
	prevStart := float64(nums[i])

	max := func(a, b float64) float64 {
		if a > b {
			return a
		}
		return b
	}

	if len(nums) == k {
		for i := 0; i < len(nums); i++ {
			total += float64(nums[i])
		}

		return total / float64(k)
	}

	for ; i <= len(nums)-k; i++ {
		for i == 0 && (j-i) < k {
			total += float64(nums[j])
			j++
		}

		if i == 0 {
			maxAvg = max(maxAvg, total/float64(k))
			continue
		}

		total -= prevStart
		total += float64(nums[i+k-1])
		prevStart = float64(nums[i])

		maxAvg = max(maxAvg, total/float64(k))
	}

	return maxAvg
}
