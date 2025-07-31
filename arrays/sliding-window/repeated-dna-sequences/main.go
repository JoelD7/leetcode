package repeated_dna_sequences

func findRepeatedDnaSequences(s string) []string {
	result := make([]string, 0)
	m := make(map[string]int)
	var seq string
	var ok bool
	var count int

	if len(s) < 10 {
		return result
	}

	for i := 0; i <= len(s)-10; i++ {
		seq = s[i : i+10]
		count, ok = m[seq]

		if ok && count == 1 {
			result = append(result, seq)
			m[seq] = count + 1
			continue
		}

		if count > 1 {
			continue
		}

		m[seq] = 1
	}

	return result
}
