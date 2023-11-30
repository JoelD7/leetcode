package three_sum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	triplets := make([][]int, 0)

	var l, r, cur, sum, lastL, lastR int

	for i := 0; i < len(nums)-1; i++ {
		cur = nums[i]

		if cur > 0 {
			break
		}

		if i > 0 && cur == nums[i-1] {
			continue
		}

		l = i + 1
		r = len(nums) - 1

		for l < r {

			sum = cur + nums[l] + nums[r]

			if sum == 0 {
				triplets = append(triplets, []int{cur, nums[l], nums[r]})

				lastL = nums[l]
				lastR = nums[r]

				for l < r && nums[l] == lastL {
					l++
				}

				for l < r && nums[r] == lastR {
					r--
				}

			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return triplets
}
