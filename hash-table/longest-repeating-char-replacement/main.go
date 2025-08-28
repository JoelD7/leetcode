package longest_repeating_char_replacement

func characterReplacement(s string, k int) int {
	//i := 0
	//j := 1
	maxLength := 0
	windowSize := k
	posByLetters := make(map[string][]int)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for index, r := range s {
		posByLetters[string(r)] = append(posByLetters[string(r)], index)
	}

	//Optimization idea: grab ONLY the letter that appears the most overall and appears most consecutively
	var length, diff int
	for _, positions := range posByLetters {
		windowSize = k
		length = 0

		for i, p := range positions {
			if i == 0 {
				length++
				continue
			}

			diff = p - positions[i-1]

			if diff == 1 {
				length++
				continue
			}

			if (diff - 1) > windowSize {
				length += windowSize
				maxLength = max(maxLength, length)
				windowSize = k
				length = 1
				continue
			}

			if (diff) > 1 {
				length += diff
				windowSize -= diff - 1
			}

			if i == len(positions)-1 {
				if windowSize > 0 {
					length += windowSize
				}
				maxLength = max(maxLength, length)
			}
		}
		maxLength = max(maxLength, length)
	}

	//"BAAAB", k = 2
	//for i < len(s)-1 && j < len(s) {
	//	if k == 0 && s[i] != s[j] {
	//		maxLength = max(maxLength, j-i)
	//		i++
	//		j = i + 1
	//		continue
	//	}
	//
	//	if s[i] != s[j] && windowSize == 0 {
	//		maxLength = max(maxLength, j-i)
	//		i++
	//		j = i + 1
	//		windowSize = k
	//		continue
	//	}
	//
	//	if s[i] != s[j] {
	//		windowSize--
	//	}
	//
	//	j++
	//}

	return maxLength
}
