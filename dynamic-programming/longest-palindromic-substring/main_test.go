package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	t.Run("Leetcode Example 1 (Multiple valid answers)", func(t *testing.T) {
		// "babad" can return "bab" or "aba". Both are accepted by LeetCode.
		result := longestPalindrome("babad")
		assert.Contains(t, []string{"bab", "aba"}, result)
	})

	t.Run("Leetcode Example 2 (Even length palindrome)", func(t *testing.T) {
		assert.Equal(t, "bb", longestPalindrome("cbbd"))
	})

	t.Run("Single character string", func(t *testing.T) {
		assert.Equal(t, "a", longestPalindrome("a"))
	})

	t.Run("Entire string is an odd length palindrome", func(t *testing.T) {
		assert.Equal(t, "racecar", longestPalindrome("racecar"))
	})

	t.Run("All identical characters", func(t *testing.T) {
		assert.Equal(t, "aaaa", longestPalindrome("aaaa"))
	})

	t.Run("Long palindrome located at the end of the string", func(t *testing.T) {
		assert.Equal(t, "deified", longestPalindrome("abcdeified"))
	})

	t.Run("Long palindrome located at the beginning of the string", func(t *testing.T) {
		assert.Equal(t, "level", longestPalindrome("levelxyz"))
	})

	t.Run("Unambiguous odd length palindrome in the middle", func(t *testing.T) {
		assert.Equal(t, "bacab", longestPalindrome("abacab"))
	})

	t.Run("String with no palindromes longer than 1 character", func(t *testing.T) {
		// For "abcdefg", any single character is a valid answer.
		result := longestPalindrome("abcdefg")
		assert.Contains(t, []string{"a", "b", "c", "d", "e", "f", "g"}, result)
	})
}
