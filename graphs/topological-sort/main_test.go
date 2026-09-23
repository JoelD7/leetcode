package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopologicalSort(t *testing.T) {
	tests := []struct {
		name     string
		graph    Graph
		start    int
		expected []int
	}{
		{
			name: "linear dependencies",
			graph: Graph{
				32: []int{21},
				21: []int{97},
				97: []int{4},
				4:  []int{53},
				53: []int{},
			},
			start:    32,
			expected: []int{32, 21, 97, 4, 53},
		},
		{
			name: "single node",
			graph: Graph{
				1: []int{},
			},
			start:    1,
			expected: []int{1},
		},
		{
			name: "DAG with multiple branches",
			// 1 points to 2 and 3. Both 2 and 3 point to 4.
			// In a standard DFS post-order reversal (visiting adjacency slices in order),
			// DFS(1) -> DFS(2) -> DFS(4) -> append 4 -> append 2 -> DFS(3) -> append 3 -> append 1.
			// Post-order: [4, 2, 3, 1]. Reversed: [1, 3, 2, 4].
			graph: Graph{
				1: []int{2, 3},
				2: []int{4},
				3: []int{4},
				4: []int{},
			},
			start:    1,
			expected: []int{1, 3, 2, 4},
		},
		{
			name: "ignores unreachable disconnected nodes",
			graph: Graph{
				1: []int{2},
				2: []int{},
				3: []int{4}, // Node 3 and 4 are not reachable from 1
				4: []int{},
			},
			start:    1,
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := topologicalSort(tt.graph, tt.start)
			assert.Equal(t, tt.expected, result)
		})
	}
}
