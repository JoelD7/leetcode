package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	t.Run("n = 1: Single pair", func(t *testing.T) {
		expected := []string{"()"}
		assert.ElementsMatch(t, expected, generateParenthesis(1))
	})

	t.Run("n = 2: Two pairs", func(t *testing.T) {
		expected := []string{"(())", "()()"}
		assert.ElementsMatch(t, expected, generateParenthesis(2))
	})

	t.Run("n = 3: Standard Leetcode example", func(t *testing.T) {
		expected := []string{
			"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()",
		}
		assert.ElementsMatch(t, expected, generateParenthesis(3))
	})

	t.Run("n = 4: Larger combination set", func(t *testing.T) {
		expected := []string{
			"(((())))", "((()()))", "((())())", "((()))()", "(()(()))",
			"(()()())", "(()())()", "(())(())", "(())()()", "()((()))",
			"()(()())", "()(())()", "()()(())", "()()()()",
		}
		assert.ElementsMatch(t, expected, generateParenthesis(4))
	})
}
