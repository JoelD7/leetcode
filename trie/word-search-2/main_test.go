package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWords(t *testing.T) {
	t.Run("Standard 4x4 board with multiple words", func(t *testing.T) {
		board := [][]byte{
			{'o', 'a', 'a', 'n'},
			{'e', 't', 'a', 'e'},
			{'i', 'h', 'k', 'r'},
			{'i', 'f', 'l', 'v'},
		}
		words := []string{"oath", "pea", "eat", "rain"}
		expected := []string{"eat", "oath"}

		assert.ElementsMatch(t, expected, findWords(board, words))
	})

	t.Run("Word cannot reuse the same cell", func(t *testing.T) {
		board := [][]byte{
			{'a', 'b'},
			{'c', 'd'},
		}
		words := []string{"abcb"}
		expected := []string{}

		assert.ElementsMatch(t, expected, findWords(board, words))
	})

	t.Run("Single cell board with matching word", func(t *testing.T) {
		board := [][]byte{
			{'a'},
		}
		words := []string{"a"}
		expected := []string{"a"}

		assert.ElementsMatch(t, expected, findWords(board, words))
	})

	t.Run("No words present in the board", func(t *testing.T) {
		board := [][]byte{
			{'x', 'y'},
			{'z', 'w'},
		}
		words := []string{"hello", "world"}
		expected := []string{}

		assert.ElementsMatch(t, expected, findWords(board, words))
	})

	t.Run("Prevent duplicate words in result output", func(t *testing.T) {
		board := [][]byte{
			{'o', 'a', 'b', 'n'},
			{'o', 't', 'a', 'e'},
			{'a', 'h', 'k', 'r'},
			{'a', 'f', 'l', 'v'},
		}
		words := []string{"oa", "oaa"}
		expected := []string{"oa", "oaa"}

		assert.ElementsMatch(t, expected, findWords(board, words))
	})

	t.Run("Word longer than available matching cells", func(t *testing.T) {
		board := [][]byte{
			{'a', 'a'},
		}
		words := []string{"aaa"}
		expected := []string{}

		assert.ElementsMatch(t, expected, findWords(board, words))
	})

	t.Run("Multiple words sharing prefix in winding path", func(t *testing.T) {
		board := [][]byte{
			{'a', 'b', 'c', 'e'},
			{'x', 'x', 'c', 'd'},
			{'x', 'x', 'b', 'a'},
		}
		words := []string{"abc", "abcd"}
		expected := []string{"abc", "abcd"}

		assert.ElementsMatch(t, expected, findWords(board, words))
	})

	t.Run("Long winding words sharing starting point in center", func(t *testing.T) {
		board := [][]byte{
			{'a', 'b', 'c'},
			{'a', 'e', 'd'},
			{'a', 'f', 'g'},
		}
		words := []string{"eaafgdcba", "eaabcdgfa"}
		expected := []string{"eaabcdgfa", "eaafgdcba"}

		assert.ElementsMatch(t, expected, findWords(board, words))
	})
}
