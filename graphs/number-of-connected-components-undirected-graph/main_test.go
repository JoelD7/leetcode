package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountComponents(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		edges := [][]int{{0, 1}, {1, 2}, {3, 4}}
		assert.Equal(t, 2, countComponents(5, edges))
	})

	t.Run("Example 2", func(t *testing.T) {
		edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}
		assert.Equal(t, 1, countComponents(5, edges))
	})

	t.Run("Test case 1", func(t *testing.T) {
		edges := [][]int{{2, 3}, {1, 2}, {1, 3}}
		assert.Equal(t, 2, countComponents(4, edges))
	})
}
