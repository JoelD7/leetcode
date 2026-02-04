package main

import (
	"math"
)

func maxSlidingWindow(nums []int, k int) []int {
	l, r := 0, k
	results := make([]int, 0)

	getMax := func(arr []int) int {
		max := math.MinInt32

		for _, val := range arr {
			if val > max {
				max = val
			}
		}

		return max
	}

	for r <= len(nums) {
		results = append(results, getMax(nums[l:r]))
		l++
		r++
	}

	return results
}
