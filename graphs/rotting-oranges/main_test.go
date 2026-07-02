package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrangesRotting(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
		assert.Equal(t, 4, orangesRotting(grid))
	})
}
