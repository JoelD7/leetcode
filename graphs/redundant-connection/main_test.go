package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRedundantConnection(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		edges := [][]int{{1, 2}, {1, 3}, {2, 3}}
		assert.Equal(t, []int{2, 3}, findRedundantConnection(edges))
	})

	t.Run("Example 2", func(t *testing.T) {
		edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
		assert.Equal(t, []int{1, 4}, findRedundantConnection(edges))
	})
}
