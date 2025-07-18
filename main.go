package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("Expected: [[-1,-1,2],[-1,0,1]], Got: %+v\n", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Printf("Expected: [[0,0,0]], Got: %+v\n", threeSum([]int{0, 0, 0, 0}))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	l := 0
	r := len(nums) - 1
	k := 0
	twoSum := 0 //Sum of the first two numbers of the triplet
	opposite := 0
	tripletMap := make(map[string]struct{})
	var tripletStr string

	result := make([][]int, 0)

	for l < r-1 {
		for ; l < r-1; l++ {
			//If the current left is the same as the previous one, we have already compared this value against the rest
			//of the array. Doing it again will produce duplicate triplets. Example: [-4,-1,-1,0,1,2]
			if l > 0 && nums[l] == nums[l-1] {
				continue
			}

			twoSum = nums[l] + nums[r]

			k = r - 1
			opposite = twoSum * -1
			for nums[k] >= opposite && k > l {
				tripletStr = fmt.Sprintf("%d%d%d", nums[l], nums[r], nums[k])
				if _, ok := tripletMap[tripletStr]; nums[k] == opposite && l != r && l != k && k != r && !ok {
					result = append(result, []int{nums[l], nums[r], nums[k]})
					tripletMap[tripletStr] = struct{}{}
					break
				}
				k--
			}
		}
		l = 0
		r--
	}

	return result
}
