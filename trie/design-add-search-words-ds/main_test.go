package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordDictionary(t *testing.T) {
	t.Run("Leetcode Example 1", func(t *testing.T) {
		wordDictionary := Constructor()
		wordDictionary.AddWord("bad")
		wordDictionary.AddWord("dad")
		wordDictionary.AddWord("mad")

		assert.False(t, wordDictionary.Search("pad")) // return False
		assert.True(t, wordDictionary.Search("bad"))  // return True
		assert.True(t, wordDictionary.Search(".ad"))  // return True
		assert.True(t, wordDictionary.Search("b.."))  // return True
	})

	t.Run("Exact Matches Only", func(t *testing.T) {
		wordDictionary := Constructor()
		wordDictionary.AddWord("hello")
		wordDictionary.AddWord("world")

		assert.True(t, wordDictionary.Search("hello"))
		assert.True(t, wordDictionary.Search("world"))
		assert.False(t, wordDictionary.Search("hell"))
		assert.False(t, wordDictionary.Search("helloo"))
		assert.False(t, wordDictionary.Search("word"))
	})

	t.Run("Wildcards and Length Extremes", func(t *testing.T) {
		wordDictionary := Constructor()
		wordDictionary.AddWord("a")
		wordDictionary.AddWord("a")

		assert.True(t, wordDictionary.Search("."))
		assert.True(t, wordDictionary.Search("a"))
		assert.False(t, wordDictionary.Search("aa"))
		assert.False(t, wordDictionary.Search("a."))
		assert.False(t, wordDictionary.Search(".a"))

		wordDictionary.AddWord("ab")
		assert.True(t, wordDictionary.Search("a."))
		assert.True(t, wordDictionary.Search(".b"))
		assert.True(t, wordDictionary.Search("ab"))
		assert.True(t, wordDictionary.Search(".."))
		assert.False(t, wordDictionary.Search("ab."))
	})

	t.Run("Searching Empty Data Structure", func(t *testing.T) {
		wordDictionary := Constructor()

		assert.False(t, wordDictionary.Search("a"))
		assert.False(t, wordDictionary.Search("."))
		assert.False(t, wordDictionary.Search(""))
	})

	t.Run("All Wildcards", func(t *testing.T) {
		wordDictionary := Constructor()
		wordDictionary.AddWord("cat")
		wordDictionary.AddWord("bat")
		wordDictionary.AddWord("rat")

		assert.True(t, wordDictionary.Search("..."))
		assert.False(t, wordDictionary.Search("...."))
		assert.False(t, wordDictionary.Search(".."))
	})
}
