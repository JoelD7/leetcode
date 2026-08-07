package main

func climbStairs(n int) int {
	steps := []int{1, 2}
	memo := make(map[int]int)
	var val int
	var ok bool

	var recurse func(n int) int
	recurse = func(n int) int {
		val, ok = memo[n]
		if ok {
			return val
		}

		nCount := 0
		if n < 0 {
			return 0
		} else if n == 0 {
			return 1
		}

		for _, step := range steps {
			nCount += recurse(n - step)
		}

		memo[n] = nCount
		return nCount
	}

	return recurse(n)
}
