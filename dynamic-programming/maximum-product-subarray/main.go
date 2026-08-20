package main

func maxProduct(nums []int) int {
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	curMax := nums[0]
	maxVal := curMax

	prefixProduct := make([]int, len(nums))
	prefixProduct[0] = nums[0]
	maxVal = max(maxVal, prefixProduct[0])

	for i := 1; i < len(nums); i++ {
		pp := prefixProduct[i-1]
		if pp == 0 {
			pp = 1
		}
		prefixProduct[i] = pp * nums[i]
		maxVal = max(maxVal, prefixProduct[i])
	}

	suffixProduct := make([]int, len(nums))
	suffixProduct[len(nums)-1] = nums[len(nums)-1]
	maxVal = max(maxVal, suffixProduct[len(nums)-1])
	for i := len(nums) - 2; i >= 0; i-- {
		sf := suffixProduct[i+1]
		if sf == 0 {
			sf = 1
		}

		suffixProduct[i] = sf * nums[i]
		maxVal = max(maxVal, suffixProduct[i])
	}

	return maxVal
}
