package three_sum

import (
	"sort"
)

func threeSumRefresh(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	sum, m, h := 0, 0, 0

	for l := 0; l < len(nums); l++ {
		m = l + 1
		h = len(nums) - 1

		if l > 0 && nums[l] == nums[l-1] {
			continue
		}

		for m < h && m < len(nums)-1 {
			sum = nums[l] + nums[m] + nums[h]
			if nums[m] == nums[m-1] && m-1 > l {
				m++
				continue
			}
			if sum == 0 {
				result = append(result, []int{nums[l], nums[m], nums[h]})
				m++
			} else if sum > 0 {
				h--
			} else {
				m++
			}
		}
	}

	return result
}
