package combinations

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	comb := make([]int, 0)

	backtrack(1, n, k, &comb, &res)
	return res
}

func backtrack(start, n, k int, comb *[]int, res *[][]int) {
	if len(*comb) == k {
		combCopy := make([]int, len(*comb))
		copy(combCopy, *comb)
		*res = append(*res, combCopy)
		return
	}

	for num := start; num <= n; num++ {
		*comb = append(*comb, num)
		backtrack(num+1, n, k, comb, res)
		combCopy := *comb
		*comb = combCopy[:len(combCopy)-1]
	}
}
