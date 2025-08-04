package three_sum_closest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	sum := func(a, b, c int) int { return a + b + c }
	abs := func(a int) int {
		if a > 0 {
			return a
		}
		return a * -1
	}

	var i, j, k, threeSum, closestTotal, curDiff int
	minDiff := math.MaxInt64

	for ; i < len(nums)-2; i++ {
		j = i + 1
		k = len(nums) - 1

		for j < k {
			threeSum = sum(nums[i], nums[j], nums[k])
			curDiff = abs(threeSum - target)
			if curDiff < minDiff {
				closestTotal = threeSum
				minDiff = curDiff
			}

			if threeSum > target {
				k--
			} else {
				j++
			}
		}
	}

	return closestTotal
}
