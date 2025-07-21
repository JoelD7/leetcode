package three_sum

import (
	"fmt"
	"sort"
)

// -1,0,1,2,-1,-4
// -4,-1,-1,0,1,2
func threeSumRefresh(nums []int) [][]int {
	sort.Ints(nums)
	l := 0
	r := len(nums) - 1
	twoSum := 0 //Sum of the first two numbers of the triplet
	opposite := 0
	tripletMap := make(map[string]struct{})
	var tripletStr string
	posByNums := make(map[int][]int)

	if nums[l] == nums[r] && nums[l] == 0 {
		return [][]int{nums[:3]}
	}

	for i, num := range nums {
		if pos, ok := posByNums[num]; ok {
			posByNums[num] = append(pos, i)
			continue
		}

		posByNums[num] = []int{i}
	}

	result := make([][]int, 0)

	for ; l < r-1; r-- {
		for ; l < r-1; l++ {
			//If the current left is the same as the previous one, we have already compared this value against the rest
			//of the array. Doing it again will produce duplicate triplets. Example: [-4,-1,-1,0,1,2]
			if l > 0 && nums[l] == nums[l-1] {
				continue
			}

			twoSum = nums[l] + nums[r]

			opposite = twoSum * -1
			pos, ok := posByNums[opposite]
			if !ok {
				continue
			}

			for _, p := range pos {
				tripletStr = fmt.Sprintf("%d%d%d", nums[l], nums[r], nums[p])
				if _, ok := tripletMap[tripletStr]; l < p && r > p && !ok {
					result = append(result, []int{nums[l], nums[r], nums[p]})
					tripletMap[tripletStr] = struct{}{}
					break
				}
			}
		}
		l = 0
	}

	return result
}
