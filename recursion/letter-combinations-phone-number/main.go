package letter_combinations_phone_number

func letterCombinations(digits string) []string {
	digitToLetters := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	res := make([]string, 0)

	var backtrack func(index int, curComb string)

	backtrack = func(index int, curComb string) {
		if len(curComb) == len(digits) {
			res = append(res, curComb)
			return
		}

		digitLetters := digitToLetters[digits[index]]

		for i := 0; i < len(digitLetters); i++ {
			backtrack(index+1, curComb+digitLetters[i])
		}
	}

	backtrack(0, "")
	return res
}
