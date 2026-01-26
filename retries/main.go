package main

func minEatingSpeed(piles []int, h int) int {
	isFeasible := func(k int) bool {
		if k == 0 {
			return false
		}
		sum := 0
		for i := 0; i < len(piles); i++ {
			sum += ((piles[i] - 1) / k) + 1
		}

		return sum <= h
	}

	var l, m, r int

	for i := 0; i < len(piles); i++ {
		if piles[i] > r {
			r = piles[i]
		}
	}

	for l <= r {
		m = l + (r-l)/2

		if isFeasible(m) {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return l
}
