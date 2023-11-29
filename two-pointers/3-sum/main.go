package three_sum

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	triplets := make([][]int, 0)

	var l, r, m, cur, lSum, rSum int
	var tripletStr string
	var ok bool
	addedTriplets := make(map[string]struct{})

	//-4 -1 -1 0 1 2
	for i := 0; i < len(nums)-1; i++ {
		cur = nums[i]
		l = i + 1
		r = len(nums) - 1
		m = (l + r) / 2

		lSum = cur + nums[l] + nums[m]
		rSum = cur + nums[r] + nums[m]

		if lSum < 0 {
			l = m
			m = r
		} else if lSum == 0 {
			tripletStr = fmt.Sprintf("%d%d%d", cur, nums[l], nums[r])

			_, ok = addedTriplets[tripletStr]
			if !ok {
				triplets = append(triplets, []int{cur, nums[l], nums[r]})

				addedTriplets[tripletStr] = struct{}{}
			}
		}

		for l < r {

			if lSum == 0 {
				tripletStr = fmt.Sprintf("%d%d%d", cur, nums[l], nums[r])

				_, ok = addedTriplets[tripletStr]
				if !ok {
					triplets = append(triplets, []int{cur, nums[l], nums[r]})

					addedTriplets[tripletStr] = struct{}{}
				}
				l++
			} else if lSum > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return triplets
}
