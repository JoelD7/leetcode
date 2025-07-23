package contains_duplicate_2

import (
	"math"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	var lastPos int
	var ok bool
	seen := map[int]int{}

	for i := 0; i < len(nums); i++ {
		lastPos, ok = seen[nums[i]]
		if ok && math.Abs(float64(i)-float64(lastPos)) <= float64(k) {
			return true
		}

		seen[nums[i]] = i
	}

	return false
}
