package longest_repeating_char_replacement

func characterReplacement(s string, k int) int {
	freq := make(map[byte]int)
	var maxLength, maxFreq, i, curLen int
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for j := 0; j < len(s); j++ {
		freq[s[j]]++
		curLen = j - i + 1

		maxFreq = max(maxFreq, freq[s[j]])
		if curLen-maxFreq > k {
			freq[s[i]]--
			i++
		}

		maxLength = max(maxLength, j-i+1)
	}

	return maxLength
}
