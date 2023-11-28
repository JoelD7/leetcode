package three_sum

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	triplets := make([][]int, 0)

	var l, r, cur, sum int
	var tripletStr string
	var ok bool
	addedTriplets := make(map[string]struct{})

	for i := 0; i < len(nums)-1; i++ {
		cur = nums[i]
		l = i + 1
		r = len(nums) - 1

		for l < r {

			sum = cur + nums[l] + nums[r]

			if sum == 0 {
				tripletStr = fmt.Sprintf("%d%d%d", cur, nums[l], nums[r])

				_, ok = addedTriplets[tripletStr]
				if !ok {
					triplets = append(triplets, []int{cur, nums[l], nums[r]})

					addedTriplets[tripletStr] = struct{}{}
				}
				l++
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return triplets
}
