package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	magazineFreq := make(map[rune]int)

	var ok bool
	var count int

	for _, s := range magazine {
		count, ok = magazineFreq[s]
		if ok {
			magazineFreq[s] = count + 1
			continue
		}

		magazineFreq[s] = 1
	}

	for _, s := range ransomNote {
		count, ok = magazineFreq[s]
		if !ok {
			return false
		}

		if count-1 == 0 {
			delete(magazineFreq, s)
			continue
		}
		magazineFreq[s] = count - 1
	}

	return true
}
