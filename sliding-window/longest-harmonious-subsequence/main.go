package longest_harmonious_subsequence

import (
	"sort"
)

func findLHS(nums []int) int {
	sort.Ints(nums)
	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	var i, newI, diff, count, largest int
	j := i + 1

	for j < len(nums) {
		diff = nums[j] - nums[i]
		if diff == 0 {
			j++
			continue
		}

		if diff == 1 && newI == 0 {
			newI = j
			j++
			count++
			continue
		}

		if diff == 1 {
			if count == 0 {
				count = j - i
			} else {
				count++
			}

			j++
			continue
		}

		if diff > 1 {
			largest = max(largest, count+1)
			count = 0
			i = newI
			newI = j
			j = i + 1
			continue
		}
	}

	return largest
}
