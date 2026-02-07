package longest_repeating_char_replacement

func characterReplacement(s string, k int) int {
	freqMap := make([]int, 26)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var maxLen, mostFreq, l, r, charsNeedToChange int

	for r < len(s) {
		freqMap[s[r]-'A']++
		mostFreq = max(mostFreq, freqMap[s[r]-'A'])

		charsNeedToChange = (r - l + 1) - mostFreq
		if charsNeedToChange > k {
			freqMap[s[l]-'A']--
			l++
		}

		maxLen = max(maxLen, r-l+1)
		r++
	}

	return maxLen
}
