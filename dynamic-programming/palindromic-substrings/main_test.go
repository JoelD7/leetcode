package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSubstrings(t *testing.T) {
	t.Run("Example 1 from Leetcode: 'abc'", func(t *testing.T) {
		assert.Equal(t, 3, countSubstrings("abc"))
	})

	t.Run("Example 2 from Leetcode: 'aaa'", func(t *testing.T) {
		assert.Equal(t, 6, countSubstrings("aaa"))
	})

	t.Run("Single character string", func(t *testing.T) {
		assert.Equal(t, 1, countSubstrings("z"))
	})

	t.Run("Two different characters", func(t *testing.T) {
		assert.Equal(t, 2, countSubstrings("ab"))
	})

	t.Run("Even length palindrome string", func(t *testing.T) {
		// "a", "b", "b", "a", "bb", "abba" -> 6
		assert.Equal(t, 6, countSubstrings("abba"))
	})

	t.Run("No palindromes longer than 1 character", func(t *testing.T) {
		assert.Equal(t, 5, countSubstrings("abcde"))
	})

	t.Run("Complex nested palindromes", func(t *testing.T) {
		// "a"(4), "b"(2), "c"(1), "aba"(2), "aca"(1), "bacab"(1), "abacaba"(1) -> 12
		assert.Equal(t, 12, countSubstrings("abacaba"))
	})
}
