package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Run("abcabcbb", func(t *testing.T) {
		assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	})

	t.Run("bbbbb", func(t *testing.T) {
		assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	})

	t.Run("pwwkew", func(t *testing.T) {
		assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
	})

	t.Run("dvdf", func(t *testing.T) {
		assert.Equal(t, 3, lengthOfLongestSubstring("dvdf"))
	})

	t.Run("Space string", func(t *testing.T) {
		assert.Equal(t, 1, lengthOfLongestSubstring(" "))
	})
}
