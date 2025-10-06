package longest_repeating_char_replacement

func characterReplacement(s string, k int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	freq := map[string]int{}
	var i, maxLength, maxFreq, curLen int

	for j := 0; j < len(s); j++ {
		curLen = j - i + 1

		if val, ok := freq[string(s[j])]; ok {
			freq[string(s[j])] = val + 1
		} else {
			freq[string(s[j])] = 1
		}

		maxFreq = max(freq[string(s[j])], maxFreq)
		if curLen-maxFreq > k {
			freq[string(s[i])]--
			i++
			continue
		}

		maxLength = max(maxLength, curLen)
	}

	return maxLength
}
