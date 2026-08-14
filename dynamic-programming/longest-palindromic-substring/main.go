package main

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	var expandFromCenter func(l, r int) string
	expandFromCenter = func(l, r int) string {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		return s[l+1 : r]
	}

	var odd, even, candidate, result string
	for i := 0; i < len(s); i++ {
		odd = expandFromCenter(i, i)
		even = expandFromCenter(i, i+1)

		if len(odd) > len(even) {
			candidate = odd
		} else {
			candidate = even
		}

		if len(candidate) > len(result) {
			result = candidate
		}
	}

	return result
}
