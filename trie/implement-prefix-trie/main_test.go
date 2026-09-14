package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	t.Run("Standard LeetCode Example", func(t *testing.T) {
		trie := Constructor()

		trie.Insert("apple")
		assert.True(t, trie.Search("apple"))
		assert.False(t, trie.Search("app"))
		assert.True(t, trie.StartsWith("app"))

		trie.Insert("app")
		assert.True(t, trie.Search("app"))
	})

	t.Run("Distinct unrelated words", func(t *testing.T) {
		trie := Constructor()

		trie.Insert("cat")
		trie.Insert("dog")

		assert.True(t, trie.Search("cat"))
		assert.True(t, trie.Search("dog"))
		assert.False(t, trie.Search("ca"))
		assert.True(t, trie.StartsWith("ca"))
		assert.False(t, trie.Search("do"))
		assert.True(t, trie.StartsWith("do"))
		assert.False(t, trie.Search("bird"))
		assert.False(t, trie.StartsWith("bir"))
	})

	t.Run("Overlapping prefixes", func(t *testing.T) {
		trie := Constructor()

		trie.Insert("hello")
		trie.Insert("hell")
		trie.Insert("heaven")
		trie.Insert("heavy")

		assert.True(t, trie.Search("hello"))
		assert.True(t, trie.Search("hell"))
		assert.False(t, trie.Search("he"))
		assert.True(t, trie.StartsWith("he"))

		assert.True(t, trie.Search("heaven"))
		assert.True(t, trie.Search("heavy"))
		assert.False(t, trie.Search("heavyweight"))
		assert.False(t, trie.StartsWith("heavyweight"))
	})

	t.Run("Single character words and prefixes", func(t *testing.T) {
		trie := Constructor()

		trie.Insert("a")
		assert.True(t, trie.Search("a"))
		assert.True(t, trie.StartsWith("a"))
		assert.False(t, trie.Search("b"))
		assert.False(t, trie.StartsWith("b"))

		trie.Insert("b")
		assert.True(t, trie.Search("b"))
		assert.True(t, trie.StartsWith("b"))
	})

	t.Run("Searching for a longer string than inserted", func(t *testing.T) {
		trie := Constructor()

		trie.Insert("car")
		assert.False(t, trie.Search("cars"))
		assert.False(t, trie.StartsWith("cars"))
		assert.True(t, trie.Search("car"))
	})
}
