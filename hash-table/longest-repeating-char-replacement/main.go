package longest_repeating_char_replacement

func characterReplacement(s string, k int) int {
	i := 0
	j := 1
	maxLength := 0
	windowSize := k

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	//"BAAAB", k = 2
	for i < len(s)-1 && j < len(s) {
		if k == 0 && s[i] != s[j] {
			maxLength = max(maxLength, j-i)
			i++
			j = i + 1
			continue
		}

		if s[i] != s[j] && windowSize == 0 {
			maxLength = max(maxLength, j-i)
			i++
			j = i + 1
			windowSize = k
			continue
		}

		if s[i] != s[j] {
			windowSize--
		}

		j++
	}

	maxLength = max(maxLength, j-i)

	//ABBB, k = 2
	i = len(s) - 1
	j = i - 1
	for i > 0 && j >= 0 {
		if k == 0 && s[i] != s[j] {
			maxLength = max(maxLength, i-j)
			i--
			j = i - 1
			continue
		}

		if s[i] != s[j] && windowSize == 0 {
			maxLength = max(maxLength, i-j)
			i--
			j = i - 1
			windowSize = k
			continue
		}

		if s[i] != s[j] {
			windowSize--
		}

		j--
	}

	return max(maxLength, i-j)
}
