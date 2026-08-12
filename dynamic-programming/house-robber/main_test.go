package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRob(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 1}
		assert.Equal(t, 4, rob(nums))
	})

	t.Run("Example 2", func(t *testing.T) {
		nums := []int{2, 7, 9, 3, 1}
		assert.Equal(t, 12, rob(nums))
	})

	t.Run("Single house", func(t *testing.T) {
		nums := []int{5}
		assert.Equal(t, 5, rob(nums))
	})

	t.Run("Two houses, second is richer", func(t *testing.T) {
		nums := []int{10, 20}
		assert.Equal(t, 20, rob(nums))
	})

	t.Run("Two houses, first is richer", func(t *testing.T) {
		nums := []int{20, 10}
		assert.Equal(t, 20, rob(nums))
	})

	t.Run("Rob first and last house", func(t *testing.T) {
		nums := []int{2, 1, 1, 2}
		assert.Equal(t, 4, rob(nums))
	})

	t.Run("All houses have the same stash", func(t *testing.T) {
		nums := []int{5, 5, 5, 5, 5}
		// Can rob 1st, 3rd, and 5th house => 5 + 5 + 5 = 15
		assert.Equal(t, 15, rob(nums))
	})

	t.Run("Zero stash everywhere", func(t *testing.T) {
		nums := []int{0, 0, 0, 0}
		assert.Equal(t, 0, rob(nums))
	})

	t.Run("Alternating highs and lows", func(t *testing.T) {
		nums := []int{10, 1, 1, 10, 1, 10}
		// Can rob 1st, 4th, and 6th => 10 + 10 + 10 = 30
		assert.Equal(t, 30, rob(nums))
	})
}
