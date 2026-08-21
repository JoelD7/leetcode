package main

import (
	"strings"
)

func wordBreak(s string, wordDict []string) bool {
	memo := make(map[string]bool)

	var recurse func(s string, wordDict []string) bool
	recurse = func(str string, wordDict []string) bool {
		if str == "" {
			return true
		}

		val, exists := memo[str]
		if exists {
			return val
		}

		var tmp string

		for _, word := range wordDict {
			if strings.HasPrefix(str, word) {
				tmp = str[len(word):]

				if recurse(tmp, wordDict) {
					memo[tmp] = true
					return true
				}
			}
		}

		memo[str] = false
		return false
	}

	return recurse(s, wordDict)
}
