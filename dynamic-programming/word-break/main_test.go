package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordBreak(t *testing.T) {
	t.Run("leetcode exact split", func(t *testing.T) {
		s := "leetcode"
		wordDict := []string{"leet", "code"}
		assert.Equal(t, true, wordBreak(s, wordDict))
	})

	t.Run("repeated dictionary words allowed", func(t *testing.T) {
		s := "applepenapple"
		wordDict := []string{"apple", "pen"}
		assert.Equal(t, true, wordBreak(s, wordDict))
	})

	t.Run("cannot break correctly", func(t *testing.T) {
		s := "catsandog"
		wordDict := []string{"cats", "dog", "sand", "and", "cat"}
		assert.Equal(t, false, wordBreak(s, wordDict))
	})

	t.Run("single character match", func(t *testing.T) {
		s := "a"
		wordDict := []string{"a"}
		assert.Equal(t, true, wordBreak(s, wordDict))
	})

	t.Run("single character mismatch", func(t *testing.T) {
		s := "a"
		wordDict := []string{"b"}
		assert.Equal(t, false, wordBreak(s, wordDict))
	})

	t.Run("complex overlapping words", func(t *testing.T) {
		s := "cars"
		wordDict := []string{"car", "ca", "rs"}
		assert.Equal(t, true, wordBreak(s, wordDict))
	})

	t.Run("unreachable end due to odd length combination", func(t *testing.T) {
		s := "aaaaaaa"
		wordDict := []string{"aaaa", "aa"}
		assert.Equal(t, false, wordBreak(s, wordDict))
	})

	t.Run("ccbb", func(t *testing.T) {
		s := "ccbb"
		wordDict := []string{"bc", "cb"}
		assert.Equal(t, false, wordBreak(s, wordDict))
	})
}
