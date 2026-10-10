package subsets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsets(t *testing.T) {
	t.Run("Example 1: Three elements", func(t *testing.T) {
		nums := []int{1, 2, 3}
		expected := [][]int{
			nil, //Leetcode treats "nil" and "{}" as the same; Go doesn't, we use "nil" in unit tests to avoid false negatives
			{1},
			{2},
			{1, 2},
			{3},
			{1, 3},
			{2, 3},
			{1, 2, 3},
		}

		// Using ElementsMatch because LeetCode allows the answer in any order
		assert.ElementsMatch(t, expected, subsets(nums))
	})

	t.Run("Example 2: Single element", func(t *testing.T) {
		nums := []int{0}
		expected := [][]int{
			nil,
			{0},
		}

		assert.ElementsMatch(t, expected, subsets(nums))
	})

	t.Run("Two elements", func(t *testing.T) {
		nums := []int{1, 2}
		expected := [][]int{
			nil,
			{1},
			{2},
			{1, 2},
		}

		assert.ElementsMatch(t, expected, subsets(nums))
	})

	t.Run("Empty array", func(t *testing.T) {
		nums := []int{}
		expected := [][]int{
			nil,
		}

		assert.ElementsMatch(t, expected, subsets(nums))
	})
}
