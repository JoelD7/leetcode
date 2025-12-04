package letter_combinations_phone_number

func letterCombinations(digits string) []string {
	var comb string
	ln := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	return backtrack(digits, comb, make([]string, 0), ln)
}

func backtrack(digits, comb string, res []string, ln map[byte][]string) []string {
	var letters []string
	if digits == "" {
		return res
	}

	for i := 0; i < len(digits); i++ {
		letters = ln[digits[i]]

		for j := 0; j < len(letters); j++ {
			comb += letters[j]
			res = backtrack(digits[i+1:], comb, res, ln)
			res = append(res, comb)
			comb = comb[:len(comb)-1]
		}
	}

	return res
}
