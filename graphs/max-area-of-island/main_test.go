package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxAreaOfIsland(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		grid := [][]int{
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}

		assert.Equal(t, 6, maxAreaOfIsland(grid))
	})

	t.Run("Example 1", func(t *testing.T) {
		grid := [][]int{{0, 0, 0, 0, 0, 0, 0, 0}}

		assert.Equal(t, 0, maxAreaOfIsland(grid))
	})

	t.Run("Test case 16", func(t *testing.T) {
		grid := [][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}
		assert.Equal(t, 4, maxAreaOfIsland(grid))
	})

}
