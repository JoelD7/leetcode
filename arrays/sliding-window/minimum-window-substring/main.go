package minimum_window_substring

import (
	"math"
)

func minWindow(s string, t string) string {
	charCount := map[byte]int{}
	targetCharsRemaining := len(t)
	minStart := 0
	minEnd := math.MaxInt32
	var i int
	var charAtStart byte

	if len(t) > len(s) {
		return ""
	}

	for j := 0; j < len(t); j++ {
		if _, ok := charCount[t[j]]; ok {
			charCount[t[j]]++
		} else {
			charCount[t[j]] = 1
		}
	}

	for j := 0; j < len(s); j++ {
		if val, ok := charCount[s[j]]; ok && val > 0 {
			targetCharsRemaining--
		}
		charCount[s[j]]--

		if targetCharsRemaining == 0 {
			for {
				charAtStart = s[i]
				if val, _ := charCount[charAtStart]; val == 0 {
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
