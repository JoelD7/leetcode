package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLadderLength(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		beginWord := "hit"
		endWord := "cog"
		wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

		assert.Equal(t, 5, ladderLength(beginWord, endWord, wordList))
	})

	t.Run("Test case 1", func(t *testing.T) {
		beginWord := "hot"
		endWord := "dog"
		wordList := []string{"hot", "dog"}

		assert.Equal(t, 0, ladderLength(beginWord, endWord, wordList))
	})

	t.Run("Test case 2", func(t *testing.T) {
		beginWord := "hot"
		endWord := "dog"
		wordList := []string{"hot", "dog", "dot"}

		assert.Equal(t, 3, ladderLength(beginWord, endWord, wordList))
	})
}
