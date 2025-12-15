package valid_palindrome

import (
	"strings"
)

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	l := 0
	r := len(s) - 1
	s = strings.ToLower(s)
	alphaNumerics := map[byte]bool{
		'a': true,
		'b': true,
		'c': true,
		'd': true,
		'e': true,
		'f': true,
		'g': true,
		'h': true,
		'i': true,
		'j': true,
		'k': true,
		'l': true,
		'm': true,
		'n': true,
		'o': true,
		'p': true,
		'q': true,
		'r': true,
		's': true,
		't': true,
		'u': true,
		'v': true,
		'w': true,
		'x': true,
		'y': true,
		'z': true,
		'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
	}

	for l < r {
		if !alphaNumerics[s[l]] {
			l++
			continue
		}

		if !alphaNumerics[s[r]] {
			r--
			continue
		}

		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}
