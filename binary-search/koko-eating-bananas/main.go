package koko_eating_bananas

import (
	"sort"
)

func minEatingSpeed(piles []int, h int) int {
	isPossible := func(speed int) bool {
		var sum int
		for _, pile := range piles {
			sum += (pile-1)/speed + 1
		}

		return sum <= h
	}

	sort.Ints(piles)

	left, right := 1, piles[len(piles)-1]
	var mid int

	for left < right {
		mid = left + (right-left)/2

		if isPossible(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
