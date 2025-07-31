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

	i := 0
	j := i + 1
	k := len(nums) - 1

	threeSum := 0
	closestTotal := 0
	minDiff := math.MaxInt64
	curDiff := 0

	for ; i < len(nums)-2; i++ {
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
