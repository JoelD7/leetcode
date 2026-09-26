package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	t.Run("Example 1: varied terrain", func(t *testing.T) {
		height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
		assert.Equal(t, 6, trap(height))
	})

	t.Run("Example 2: deep basin", func(t *testing.T) {
		height := []int{4, 2, 0, 3, 2, 5}
		assert.Equal(t, 9, trap(height))
	})

	t.Run("Empty map traps 0 water", func(t *testing.T) {
		height := []int{}
		assert.Equal(t, 0, trap(height))
	})

	t.Run("Less than three bars cannot trap water", func(t *testing.T) {
		height := []int{2, 1}
		assert.Equal(t, 0, trap(height))
	})

	t.Run("Monotonically increasing terrain", func(t *testing.T) {
		height := []int{1, 2, 3, 4, 5}
		assert.Equal(t, 0, trap(height))
	})

	t.Run("Monotonically decreasing terrain", func(t *testing.T) {
		height := []int{5, 4, 3, 2, 1}
		assert.Equal(t, 0, trap(height))
	})

	t.Run("Flat terrain", func(t *testing.T) {
		height := []int{2, 2, 2, 2, 2}
		assert.Equal(t, 0, trap(height))
	})

	t.Run("Single wide and deep basin", func(t *testing.T) {
		height := []int{5, 0, 0, 0, 5}
		assert.Equal(t, 15, trap(height))
	})

	t.Run("Multiple basins of the same height", func(t *testing.T) {
		height := []int{3, 0, 3, 0, 3}
		assert.Equal(t, 6, trap(height))
	})

	t.Run("Asymmetrical basins", func(t *testing.T) {
		height := []int{5, 1, 3, 2, 4}
		assert.Equal(t, 7, trap(height))
	})
}
