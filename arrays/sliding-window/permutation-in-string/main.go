package main

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	var start, end int
	freqS1 := make([]byte, 26)
	freqS2 := make([]byte, 26)
	areArraysEqual := func(a, b []byte) bool {
		for i := 0; i < 26; i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(s1); i++ {
		freqS1[s1[i]-'a']++
	}

	for end < len(s2) {
		freqS2[s2[end]-'a']++

		if end-start+1 == len(s1) && areArraysEqual(freqS1, freqS2) {
			return true
		}

		if end-start+1 < len(s1) {
			end++
		} else {
			freqS2[s2[start]-'a']--
			start++
			end++
		}

	}

	return false
}
