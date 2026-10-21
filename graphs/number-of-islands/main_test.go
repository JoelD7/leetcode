package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumIslands(t *testing.T) {
	t.Run("Example 1 - Single large connected island", func(t *testing.T) {
		grid := [][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}
		assert.Equal(t, 1, numIslands(grid))
	})

	t.Run("Example 2 - Three separate islands", func(t *testing.T) {
		grid := [][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		}
		assert.Equal(t, 3, numIslands(grid))
	})

	t.Run("Empty grid", func(t *testing.T) {
		grid := [][]byte{}
		assert.Equal(t, 0, numIslands(grid))
	})

	t.Run("All water grid", func(t *testing.T) {
		grid := [][]byte{
			{'0', '0', '0'},
			{'0', '0', '0'},
			{'0', '0', '0'},
		}
		assert.Equal(t, 0, numIslands(grid))
	})

	t.Run("All land grid", func(t *testing.T) {
		grid := [][]byte{
			{'1', '1'},
			{'1', '1'},
		}
		assert.Equal(t, 1, numIslands(grid))
	})

	t.Run("Diagonal islands not connected", func(t *testing.T) {
		grid := [][]byte{
			{'1', '0', '1'},
			{'0', '1', '0'},
			{'1', '0', '1'},
		}
		// Diagonal connections do not count as contiguous lands.
		assert.Equal(t, 5, numIslands(grid))
	})

	t.Run("Single row grid", func(t *testing.T) {
		grid := [][]byte{
			{'1', '0', '1', '1', '0', '1'},
		}
		assert.Equal(t, 3, numIslands(grid))
	})

	t.Run("Single column grid", func(t *testing.T) {
		grid := [][]byte{
			{'1'},
			{'0'},
			{'1'},
			{'1'},
		}
		assert.Equal(t, 2, numIslands(grid))
	})
}
