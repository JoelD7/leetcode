package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrangesRotting(t *testing.T) {
	t.Run("Standard case requiring multiple minutes", func(t *testing.T) {
		grid := [][]int{
			{2, 1, 1},
			{1, 1, 0},
			{0, 1, 1},
		}
		assert.Equal(t, 4, orangesRotting(grid))
	})

	t.Run("Impossible to rot all fresh oranges due to isolation", func(t *testing.T) {
		grid := [][]int{
			{2, 1, 1},
			{0, 1, 1},
			{1, 0, 1},
		}
		assert.Equal(t, -1, orangesRotting(grid))
	})

	t.Run("Already complete with zero fresh oranges initially", func(t *testing.T) {
		grid := [][]int{
			{0, 2},
		}
		assert.Equal(t, 0, orangesRotting(grid))
	})

	t.Run("Empty grid with absolutely no oranges", func(t *testing.T) {
		grid := [][]int{
			{0, 0, 0},
			{0, 0, 0},
		}
		assert.Equal(t, 0, orangesRotting(grid))
	})

	t.Run("Multiple rotten oranges starting simultaneously", func(t *testing.T) {
		grid := [][]int{
			{2, 1},
			{1, 2},
		}
		assert.Equal(t, 1, orangesRotting(grid))
	})

	t.Run("Large spiral grid", func(t *testing.T) {
		grid := [][]int{
			{2, 0, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 0, 1, 0, 0, 0, 0, 0, 0, 1},
			{1, 0, 1, 0, 1, 1, 1, 1, 0, 1},
			{1, 0, 1, 0, 1, 0, 0, 1, 0, 1},
			{1, 0, 1, 0, 1, 0, 0, 1, 0, 1},
			{1, 0, 1, 0, 1, 1, 0, 1, 0, 1},
			{1, 0, 1, 0, 0, 0, 0, 1, 0, 1},
			{1, 0, 1, 1, 1, 1, 1, 1, 0, 1},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		}
		// The rot travels in a long spiral path taking exactly 58 minutes to reach the center
		assert.Equal(t, 58, orangesRotting(grid))
	})
}
