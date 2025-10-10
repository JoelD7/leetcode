package longest_repeating_char_replacement

func characterReplacement(s string, k int) int {
	freq := map[string]int{}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var curLen, i, mostFreq, maxLength int
	var ok bool

	for j := 0; j < len(s); j++ {
		curLen = j - i + 1

		if _, ok = freq[string(s[j])]; ok {
			freq[string(s[j])]++
		} else {
			freq[string(s[j])] = 1
		}

		mostFreq = max(mostFreq, freq[string(s[j])])
		if curLen-mostFreq > k {
			freq[string(s[i])]--
			i++
		}

		maxLength = max(maxLength, j-i+1)
	}

	return maxLength
}
