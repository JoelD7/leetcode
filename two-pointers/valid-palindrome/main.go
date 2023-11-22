package valid_palindrome

import "strings"

func isPalindrome(s string) bool {
	s = normalize(s)
	if s == "" {
		return true
	}

	l := 0
	r := len(s) - 1

	for l < r {
		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}

func normalize(s string) string {
	s = strings.ToLower(s)

	var result string

	for _, c := range s {
		if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || (c >= '0' && c <= '9') {
			result += string(c)
		}
	}

	return result
}
