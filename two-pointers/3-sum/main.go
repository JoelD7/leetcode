package three_sum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)

	var k, j, sum int

	for i := 0; i < len(nums)-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		k = i + 1
		j = len(nums) - 1

		for k < j {
			sum = nums[i] + nums[k] + nums[j]

			if sum == 0 {
				res = append(res, []int{nums[i], nums[k], nums[j]})
				k++

				for nums[k] == nums[k-1] && k < j {
					k++
				}

				continue
			}

			if sum < 0 {
				k++
				continue
			}

			if sum > 0 {
				j--
				continue
			}
		}
	}

	return res
}
