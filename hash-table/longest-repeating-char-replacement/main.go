package longest_repeating_char_replacement

func characterReplacement(s string, k int) int {
	i := 0
	j := 0
	maxLength := 0
	maxFreq := 0
	charsNeedChange := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	freq := map[string]int{}

	for j < len(s) {
		valueI := string(s[i])
		valueJ := string(s[j])

		_, ok := freq[valueJ]
		if ok {
			freq[valueJ]++
		} else {
			freq[valueJ] = 1
		}

		maxFreq = max(maxFreq, freq[valueJ])

		charsNeedChange = (j - i + 1) - maxFreq
		if charsNeedChange > k {
			i++ //window shrinks
			freq[valueI]--
		} else {
			maxLength = max(maxLength, j-i+1)
		}
		j++ //window increases
	}

	return maxLength
}
