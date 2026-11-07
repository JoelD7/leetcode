package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidTree(t *testing.T) {
	t.Run("Valid tree with 5 nodes", func(t *testing.T) {
		n := 5
		edges := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}
		assert.True(t, validTree(n, edges))
	})

	t.Run("Graph with cycle and 5 nodes", func(t *testing.T) {
		n := 5
		edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}
		assert.False(t, validTree(n, edges))
	})

	t.Run("Disconnected graph with 4 nodes", func(t *testing.T) {
		n := 4
		edges := [][]int{{0, 1}, {2, 3}}
		assert.False(t, validTree(n, edges))
	})

	t.Run("Single node graph", func(t *testing.T) {
		n := 1
		edges := [][]int{}
		assert.True(t, validTree(n, edges))
	})
}
