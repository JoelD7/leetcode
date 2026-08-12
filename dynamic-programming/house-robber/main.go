package main

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var prev1, prev2 int
	for _, num := range nums {
		tmp := prev1
		prev1 = max(prev2+num, prev1)
		prev2 = tmp
	}

	return prev1
}
