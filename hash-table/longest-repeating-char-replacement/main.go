package longest_repeating_char_replacement

func characterReplacement(s string, k int) int {
	freq := make(map[byte]int)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var i, maxFreq, charsNeedChange, curLen, maxLen int

	for j := 0; j < len(s); j++ {
		curLen = j - i + 1

		freq[s[j]]++
		maxFreq = max(maxFreq, freq[s[j]])
		charsNeedChange = curLen - maxFreq

		if charsNeedChange > k {
			freq[s[i]]--
			i++
		}

		maxLen = max(maxLen, j-i+1)
	}

	return maxLen
}
