package missing_number

import (
	"sort"
)

func missingNumber(nums []int) int {
	sort.Ints(nums)
	max := nums[len(nums)-1]
	expected := 0

	for _, num := range nums {
		if num != expected {
			return expected
		}
		expected++
	}

	if expected > max {
		return expected
	}

	return 0
}
