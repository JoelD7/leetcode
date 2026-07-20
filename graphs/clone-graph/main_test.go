package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloneGraph(t *testing.T) {
	tests := []struct {
		name    string
		adjList [][]int
	}{
		{
			name:    "Example 1: 4 Nodes connected in a square",
			adjList: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
		},
		{
			name:    "Example 2: 1 Node with no neighbors",
			adjList: [][]int{{}}, // Represents a single node with Val=1 and empty neighbors
		},
		{
			name:    "Example 3: Empty Graph",
			adjList: [][]int{}, // Represents a nil graph
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := buildGraph(tt.adjList)

			// Call the function being tested
			clone := cloneGraph(original)

			// Validate the results using testify/assert
			validateDeepCopy(t, original, clone)
		})
	}
}

// buildGraph constructs a graph from an adjacency list and returns the root node.
func buildGraph(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}

	nodes := make(map[int]*Node)

	// First pass: create all nodes
	for i := 1; i <= len(adjList); i++ {
		nodes[i] = &Node{Val: i}
	}

	// Second pass: wire up the neighbors
	for i, neighbors := range adjList {
		for _, neighborVal := range neighbors {
			nodes[i+1].Neighbors = append(nodes[i+1].Neighbors, nodes[neighborVal])
		}
	}

	return nodes[1]
}

// validateDeepCopy recursively ensures the cloned graph is a perfect deep copy.
func validateDeepCopy(t *testing.T, original *Node, clone *Node) {
	// Handle Example 3 (Empty Graph)
	if original == nil {
		assert.Nil(t, clone, "Clone should be nil when the original is nil")
		return
	}
	if !assert.NotNil(t, clone, "Clone should not be nil when original exists") {
		return
	}

	// Maps an original node's memory address to its cloned counterpart
	visitedMap := make(map[*Node]*Node)

	var dfs func(o, c *Node)
	dfs = func(o, c *Node) {
		// 1. Validate it is actually a deep copy (different memory addresses)
		assert.NotSame(t, o, c, "Nodes must have different memory addresses (not a deep copy)")

		// 2. Validate Values and number of neighbors match
		assert.Equal(t, o.Val, c.Val, "Node values must match")
		assert.Equal(t, len(o.Neighbors), len(c.Neighbors), "Number of neighbors must match")

		// Mark as visited by keeping track of the mapping
		visitedMap[o] = c

		for i := 0; i < len(o.Neighbors); i++ {
			oNeighbor := o.Neighbors[i]
			cNeighbor := c.Neighbors[i]

			if mappedClone, exists := visitedMap[oNeighbor]; !exists {
				// Neighbor not visited yet, traverse deeper
				dfs(oNeighbor, cNeighbor)
			} else {
				// 3. Topology validation: The clone's cyclic reference MUST point
				// to the exact same memory address as the previously cloned node.
				assert.Same(t, mappedClone, cNeighbor, "Clone graph topology does not mirror original cyclic dependencies")
			}
		}
	}

	dfs(original, clone)
}
