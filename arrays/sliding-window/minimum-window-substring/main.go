package minimum_window_substring

import (
	"math"
)

func minWindow(s string, t string) string {
	m := len(s)
	n := len(t)
	tFreq := map[string]int{}
	tsUsed := n
	minLength := math.MaxInt32
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var i int

	if n > m {
		return ""
	}

	for j := 0; j < n; j++ {
		if _, ok := tFreq[string(t[j])]; ok {
			tFreq[string(t[j])]++
		} else {
			tFreq[string(t[j])] = 1
		}
	}

	for j := 0; j < len(s); j++ {
		if val, ok := tFreq[string(s[j])]; ok && val > 0 {
			tFreq[string(s[j])]-- // I shouldn't be updating this map because I need it for other iterations.
			tsUsed--
		}

		//Input: s = "ADOBEC ODEBANC", t = "ABC"
		//Output: "BANC"
		if tsUsed == 0 {
			minLength = min(minLength, j-i+1)
			tsUsed = n
			i++
		}
	}
}
