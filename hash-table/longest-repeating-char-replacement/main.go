package longest_repeating_char_replacement

func characterReplacement(s string, k int) int {
	var i, maxLength, curLen, maxFreq int
	var ok bool
	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}
	freq := make(map[string]int)

	for j := 0; j < len(s); j++ {
		curLen = j - i + 1

		_, ok = freq[string(s[j])]
		if ok {
			freq[string(s[j])]++
		} else {
			freq[string(s[j])] = 1
		}

		maxFreq = max(maxFreq, freq[string(s[j])])

		if curLen-maxFreq > k {
			freq[string(s[i])]--
			i++
		}
		maxLength = max(maxLength, j-i+1)
	}

	return maxLength
}
