package main

func lengthOfLongestSubstring(s string) int {
	table := make(map[byte]bool)
	l, r, maxLength := 0, 0, 0
	isDuplicated := false

	for r < len(s) {
		_, isDuplicated = table[s[r]]
		if isDuplicated {
			l++
			r = l
			table = map[byte]bool{}
		} else {
			if (r-l)+1 > maxLength {
				maxLength = r - l + 1
			}
			table[s[r]] = true
			r++
		}

	}

	return maxLength
}
