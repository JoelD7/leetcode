package minimum_window_substring

import (
	"math"
)

func minWindow(s string, t string) string {
	m := len(s)
	n := len(t)
	charCount := map[string]int{}
	targetCharsRemaining := n
	minWindowArr := []int{0, math.MaxInt32}
	var i int
	var charAtStart string

	if n > m {
		return ""
	}

	for j := 0; j < n; j++ {
		if _, ok := charCount[string(t[j])]; ok {
			charCount[string(t[j])]++
		} else {
			charCount[string(t[j])] = 1
		}
	}

	for j := 0; j < len(s); j++ {
		if val, ok := charCount[string(s[j])]; ok && val > 0 {
			targetCharsRemaining--
		}
		charCount[string(s[j])]--

		if targetCharsRemaining == 0 {
			for {
				charAtStart = string(s[i])
				if val, _ := charCount[charAtStart]; val == 0 {
					break
				}
				charCount[charAtStart]++
				i++
			}

			if j-i < minWindowArr[1]-minWindowArr[0] {
				minWindowArr = []int{i, j}
			}
			charCount[string(s[i])]++
			targetCharsRemaining++
			i++
		}
	}

	if minWindowArr[1] > len(s) {
		return ""
	}

	return s[minWindowArr[0] : minWindowArr[1]+1]
}
