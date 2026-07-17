package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBFSGraphs(t *testing.T) {

	t.Run("EmptyGraph", func(t *testing.T) {
		adjList := [][]int{}

		traversal, levels := bfsGraphs(adjList)

		assert.Empty(t, traversal, "Expected an empty slice for an empty graph")
		assert.Equal(t, 0, levels)
	})

	t.Run("SingleIsolatedNode", func(t *testing.T) {
		adjList := [][]int{{}}
		expected := []int{0}

		traversal, levels := bfsGraphs(adjList)

		assert.Equal(t, expected, traversal, "Failed on single isolated node")
		assert.Equal(t, 1, levels)
	})

	t.Run("ScatteredTraversal", func(t *testing.T) {
		// Node zero visits node two, node two visits node one
		adjList := [][]int{{2}, {}, {1}}
		expected := []int{0, 2, 1}

		traversal, levels := bfsGraphs(adjList)

		assert.Equal(t, expected, traversal, "Failed on scattered path")
		assert.Equal(t, 3, levels)
	})

	t.Run("StarGraphMixedOrder", func(t *testing.T) {
		// Node zero connects to three and one. Three connects to nothing, one connects to two.
		adjList := [][]int{{3, 1}, {2}, {}, {}}
		expected := []int{0, 3, 1, 2}

		traversal, levels := bfsGraphs(adjList)

		assert.Equal(t, expected, traversal, "Failed on mixed breadth traversal")
		assert.Equal(t, 3, levels)
	})

	t.Run("ComplexGraphWithCycles", func(t *testing.T) {
		// Graph structure where traversal jumps back and forth:
		// Node zero connects to four and two
		// Node four points back to zero
		// Node two connects to one
		// Node one connects to three
		adjList := [][]int{
			{4, 2},
			{3},
			{1},
			{},
			{0},
		}
		expected := []int{0, 4, 2, 1, 3}

		traversal, levels := bfsGraphs(adjList)

		assert.Equal(t, expected, traversal, "Failed to handle cycles and prevent infinite loops")
		assert.Equal(t, 4, levels)
	})

	t.Run("DisconnectedComponents", func(t *testing.T) {
		// Node zero only connects to node five. Nodes two and four are unreachable.
		adjList := [][]int{{5}, {0}, {4}, {}, {}, {}}
		expected := []int{0, 5}

		traversal, levels := bfsGraphs(adjList)

		assert.Equal(t, expected, traversal, "Failed on disconnected components")
		assert.Equal(t, 2, levels)
	})
}
