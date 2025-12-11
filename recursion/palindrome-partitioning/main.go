package palindrome_partitioning

func partition(s string) [][]string {
	res := make([][]string, 0)
	comb := make([]string, 0)

	var backtrack func(start int, str string)

	backtrack = func(start int, str string) {
		if str == "" {
			res = append(res, append([]string{}, comb...))
			return
		}

		for i := 0; i < len(str); i++ {
			if isPalindrome(str[:i+1]) {
				comb = append(comb, str[:i+1])
				backtrack(i+1, str[i+1:])
				comb = comb[:len(comb)-1]
			}
		}
	}

	backtrack(0, s)
	return res
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
