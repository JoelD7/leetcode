package longest_consecutive_sequence

import (
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	maxFreq := 0
	count := 1
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i]-1 {
			count++
			continue
		}

		if nums[i-1] == nums[i] {
			continue
		}

		maxFreq = max(maxFreq, count)
		count = 1
	}

	maxFreq = max(maxFreq, count)

	return maxFreq
}
