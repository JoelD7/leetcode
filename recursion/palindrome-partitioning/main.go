package palindrome_partitioning

func partition(s string) [][]string {
	comb := make([]string, 0)
	res := make([][]string, 0)

	return backtrack(s, comb, res)
}

func backtrack(s string, comb []string, res [][]string) [][]string {
	if s == "" {
		res = append(res, append([]string(nil), comb...))
		return res
	}

	for i := 1; i <= len(s); i++ {
		if isPalindrome(s[:i]) {
			comb = append(comb, s[:i])
			res = backtrack(s[i:], comb, res)
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
