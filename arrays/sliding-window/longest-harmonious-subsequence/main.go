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
	updateLargest := func(largest, count int) int {
		if count > 0 {
			return max(largest, count+1)
		}

		return largest
	}

	var i, newI, dif, count, largest int
	j := i + 1

	for j < len(nums) {
		dif = nums[j] - nums[i]
		if dif == 0 {
			j++
			continue
		}

		if dif == 1 || dif == -1 {
			if nums[j-1] != nums[j] {
				newI = j
			}

			count = j - i
			j++
			continue
		}

		if (dif > 1 || dif < -1) && newI == 0 {
			newI = j
			i++
			j++
			continue
		}

		if dif > 1 || dif < -1 {
			largest = updateLargest(largest, count)
			count = 0
			i = newI
			newI = j
			j = i + 1
			continue
		}
	}

	largest = updateLargest(largest, count)

	return largest
}
