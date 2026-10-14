package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCombinations(t *testing.T) {
	t.Run("digits = '23' (Two digits)", func(t *testing.T) {
		digits := "23"
		expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
		assert.ElementsMatch(t, expected, letterCombinations(digits))
	})

	t.Run("digits = '' (Empty string)", func(t *testing.T) {
		digits := ""
		expected := []string{""}
		assert.ElementsMatch(t, expected, letterCombinations(digits))
	})

	t.Run("digits = '2' (Single digit)", func(t *testing.T) {
		digits := "2"
		expected := []string{"a", "b", "c"}
		assert.ElementsMatch(t, expected, letterCombinations(digits))
	})

	t.Run("digits = '7' (Single digit with four letters)", func(t *testing.T) {
		digits := "7"
		expected := []string{"p", "q", "r", "s"}
		assert.ElementsMatch(t, expected, letterCombinations(digits))
	})

	t.Run("digits = '9' (Another single digit with four letters)", func(t *testing.T) {
		digits := "9"
		expected := []string{"w", "x", "y", "z"}
		assert.ElementsMatch(t, expected, letterCombinations(digits))
	})

	t.Run("digits = '22' (Repeating digits)", func(t *testing.T) {
		digits := "22"
		expected := []string{"aa", "ab", "ac", "ba", "bb", "bc", "ca", "cb", "cc"}
		assert.ElementsMatch(t, expected, letterCombinations(digits))
	})
}
