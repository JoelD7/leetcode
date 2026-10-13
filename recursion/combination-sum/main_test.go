package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	t.Run("Standard Leetcode example 1", func(t *testing.T) {
		candidates := []int{2, 3, 6, 7}
		target := 7
		expected := [][]int{
			{2, 2, 3},
			{7},
		}
		assert.ElementsMatch(t, expected, combinationSum(candidates, target))
	})

	t.Run("Standard Leetcode example 2 - multiple combinations", func(t *testing.T) {
		candidates := []int{2, 3, 5}
		target := 8
		expected := [][]int{
			{2, 2, 2, 2},
			{2, 3, 3},
			{3, 5},
		}
		assert.ElementsMatch(t, expected, combinationSum(candidates, target))
	})

	t.Run("No valid combination possible", func(t *testing.T) {
		candidates := []int{2}
		target := 1
		expected := [][]int{}
		assert.ElementsMatch(t, expected, combinationSum(candidates, target))
	})

	t.Run("Single candidate used multiple times", func(t *testing.T) {
		candidates := []int{1}
		target := 3
		expected := [][]int{
			{1, 1, 1},
		}
		assert.ElementsMatch(t, expected, combinationSum(candidates, target))
	})

	t.Run("Larger target with distinct candidates", func(t *testing.T) {
		candidates := []int{2, 7, 3, 4}
		target := 7
		expected := [][]int{
			{7},
			{3, 4},
			{2, 2, 3},
		}
		assert.ElementsMatch(t, expected, combinationSum(candidates, target))
	})
}
