package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDFS(t *testing.T) {
	tests := []struct {
		name     string
		graph    Graph
		start    int
		expected []int
	}{
		{
			name: "Single node, no edges",
			graph: Graph{
				1: {},
			},
			start:    1,
			expected: []int{1},
		},
		{
			name: "Linear path",
			graph: Graph{
				1: {2},
				2: {3},
				3: {4},
				4: {},
			},
			start:    1,
			expected: []int{1, 2, 3, 4},
		},
		{
			name: "Simple tree",
			// Graph shape:
			//      1
			//     / \
			//    2   3
			//   / \
			//  4   5
			graph: Graph{
				1: {2, 3},
				2: {4, 5},
				3: {},
				4: {},
				5: {},
			},
			start:    1,
			expected: []int{1, 2, 4, 5, 3},
		},
		{
			name: "Graph with a simple cycle",
			graph: Graph{
				1: {2},
				2: {3},
				3: {1}, // Cycle back to start
			},
			start:    1,
			expected: []int{1, 2, 3},
		},
		{
			name: "Graph with multiple cycles and cross edges",
			graph: Graph{
				1: {2, 3},
				2: {4},
				3: {4},
				4: {1, 5},
				5: {},
			},
			start: 1,
			// Trace:
			// 1 visits 2 -> 2 visits 4 -> 4 sees 1 (skip), visits 5 -> 5 is done.
			// Backtrack to 1 -> 1 visits 3 -> 3 sees 4 (skip, already visited).
			expected: []int{1, 2, 4, 5, 3},
		},
		{
			name: "Disconnected components",
			graph: Graph{
				1: {2},
				2: {},
				3: {4}, // 3 and 4 are unreachable from 1
				4: {},
			},
			start:    1,
			expected: []int{1, 2}, // Should only visit the connected component
		},
		{
			name:     "Start node not in graph",
			graph:    Graph{},
			start:    99,
			expected: []int{99}, // Treats the start node as a valid isolated node
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := DFS(tc.graph, tc.start)
			assert.Equal(t, tc.expected, actual, "DFS traversal order did not match expected output")
		})
	}
}
