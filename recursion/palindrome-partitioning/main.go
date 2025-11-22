package palindrome_partitioning

func partition(s string) [][]string {
	comb := make([]string, 0)
	res := make([][]string, 0)

	return backtrack(0, s, comb, res)
}

func backtrack(start int, s string, comb []string, res [][]string) [][]string {
	if start == len(s) {
		c := make([]string, len(comb))
		copy(c, comb)
		res = append(res, c)
		return res
	}

	for i := start; i < len(s); i++ {
		if isPalindrome(s[:i+1]) {
			comb = append(comb, s[:i+1])
			res = backtrack(i+1, s, comb, res)
			comb = comb[:len(comb)-1]
		}
	}

	return res
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
