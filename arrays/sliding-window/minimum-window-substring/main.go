package minimum_window_substring

import (
	"math"
)

func minWindow(s string, t string) string {
	charCount := make(map[byte]int)
	i, minStart, minEnd := 0, 0, math.MaxInt32
	targetCharsRemaining := len(t)
	var charAtStart byte

	if len(t) > len(s) {
		return ""
	}

	for j := 0; j < len(t); j++ {
		charCount[t[j]]++
	}

	for j := 0; j < len(s); j++ {
		if charCount[s[j]] > 0 {
			targetCharsRemaining--
		}
		charCount[s[j]]--

		if targetCharsRemaining == 0 {
			for {
				charAtStart = s[i]
				if charCount[charAtStart] == 0 {
					break
				}
				charCount[charAtStart]++
				i++
			}

			if j-i < minEnd-minStart {
				minStart = i
				minEnd = j
			}

			charCount[s[i]]++
			targetCharsRemaining++
			i++
		}
	}

	if minEnd > len(s) {
		return ""
	}

	return s[minStart : minEnd+1]
}
