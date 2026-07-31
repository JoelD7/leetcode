package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIslandsAndTreasure(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		grid := [][]int{{2147483647, -1, 0, 2147483647}, {2147483647, 2147483647, 2147483647, -1}, {2147483647, -1, 2147483647, -1}, {0, -1, 2147483647, 2147483647}}
		expected := [][]int{{3, -1, 0, 1}, {2, 2, 1, -1}, {1, -1, 2, -1}, {0, -1, 3, 4}}

		islandsAndTreasure(grid)
		assert.Equal(t, expected, grid)

	})
}
