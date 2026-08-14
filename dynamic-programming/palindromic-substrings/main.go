package main

func countSubstrings(s string) int {
	var count int

	var expandFromMiddle func(l, r int)
	expandFromMiddle = func(l, r int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
			count++
		}
	}

	for i := 0; i < len(s); i++ {
		expandFromMiddle(i, i)
		expandFromMiddle(i, i+1)
	}

	return count
}
