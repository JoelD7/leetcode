package longest_consecutive_sequence

import "sort"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	sort.Ints(nums)

	prev := 0
	cur := 0
	max := 0
	count := 1

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			continue
		}

		prev = nums[i-1]
		cur = nums[i]

		if cur == prev {
			continue
		}

		if cur == prev+1 {
			count++
			continue
		}

		if count > max {
			max = count
		}

		count = 1
	}

	if count > max {
		max = count
	}

	return max
}
