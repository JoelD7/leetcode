package combinations

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	comb := make([]int, 0)

	res = backtrack(1, n, k, res, comb)
	return res
}

func backtrack(start, n, k int, res [][]int, comb []int) [][]int {
	if len(comb) == k {
		c := make([]int, len(comb))
		copy(c, comb)
		res = append(res, c)
		return res
	}

	for i := start; i <= n; i++ {
		comb = append(comb, i)
		res = backtrack(i+1, n, k, res, comb)
		comb = comb[:len(comb)-1]
	}

	return res
}
