package letter_combinations_phone_number

func letterCombinations(digits string) []string {
	ln := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}

	res := make([]string, 0)

	var backtrack func(index int, digits, comb string)

	backtrack = func(index int, digits, comb string) {
		if index == len(digits) {
			res = append(res, comb)
			return
		}

		letters := ln[digits[index]]

		for i := 0; i < len(letters); i++ {
			backtrack(index+1, digits, comb+string(letters[i]))
		}
	}

	backtrack(0, digits, "")
	return res
}
