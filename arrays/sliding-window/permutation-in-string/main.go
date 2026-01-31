package main

func checkInclusion(s1 string, s2 string) bool {
	//assumption
	if len(s2) < len(s1) {
		return false
	}

	createMapCopy := func(mapToCopy map[byte]int) map[byte]int {
		mapCopy := make(map[byte]int)

		for k, v := range mapToCopy {
			mapCopy[k] = v
		}

		return mapCopy
	}

	l, r, freq := 0, 0, 0
	exists := false
	freqTable := make(map[byte]int)
	freqTableCopy := make(map[byte]int)

	for i := 0; i < len(s1); i++ {
		freqTable[s1[i]]++
		freqTableCopy[s1[i]]++
	}

	//not sure about this condition
	for r < len(s2) {
		freq, exists = freqTable[s2[r]]
		if !exists {
			r = l + 1
			l = r
			freqTable = createMapCopy(freqTableCopy)
			continue
		}

		if freq == 1 {
			//We have exhausted all frequencies of this character
			delete(freqTable, s2[r])
		} else {
			freqTable[s2[r]]--
		}

		if r-l+1 == len(s1) {
			return true
		}

		r++
	}

	return false
}
