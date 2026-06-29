package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumIslands(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		grid := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
		assert.Equal(t, 1, numIslands(grid))
	})

	t.Run("Example 2", func(t *testing.T) {
		grid := [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}
		assert.Equal(t, 3, numIslands(grid))
	})
}
