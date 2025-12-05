package combinations

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	comb := make([]int, 0)

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(comb) == k {
			res = append(res, append([]int(nil), comb...))
			return
		}

		for i := start; i <= n; i++ {
			comb = append(comb, i)
			backtrack(i + 1)
			comb = comb[:len(comb)-1]
		}
	}

	backtrack(1)
	return res
}
