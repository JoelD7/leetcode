package copy_list_with_random_pointer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyRandomList(t *testing.T) {
	tests := []struct {
		name string
		data [][]interface{}
	}{
		{
			name: "Example 1",
			data: [][]interface{}{
				{7, nil},
				{13, 0},
				{11, 4},
				{10, 2},
				{1, 0},
			},
		},
		{
			name: "Example 2",
			data: [][]interface{}{
				{1, 1},
				{2, 1},
			},
		},
		{
			name: "Example 3",
			data: [][]interface{}{
				{3, nil},
				{3, 0},
				{3, nil},
			},
		},
		{
			name: "Empty List",
			data: [][]interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := buildList(tt.data)

			// Call the function you are testing
			copied := copyRandomList(original)

			// Run the deep copy validation
			validateCopy(t, original, copied)
		})
	}
}

// buildList is a helper function to construct the linked list from a 2D slice
// matching the LeetCode input format: [][val, random_index]
func buildList(data [][]interface{}) *Node {
	if len(data) == 0 {
		return nil
	}

	nodes := make([]*Node, len(data))

	// First pass: create all the nodes with their values
	for i, d := range data {
		nodes[i] = &Node{Val: d[0].(int)}
	}

	// Second pass: wire up the Next and Random pointers
	for i, d := range data {
		if i < len(data)-1 {
			nodes[i].Next = nodes[i+1]
		}
		if d[1] != nil {
			nodes[i].Random = nodes[d[1].(int)]
		}
	}

	return nodes[0]
}

// validateCopy traverses both lists and asserts that the copied list
// is a true deep copy of the original list.
func validateCopy(t *testing.T, original *Node, copied *Node) {
	if original == nil {
		assert.Nil(t, copied, "Copied list should be nil if original is nil")
		return
	}

	// Maps to store the index position of each node reference
	origMap := make(map[*Node]int)
	copyMap := make(map[*Node]int)

	// Map original nodes to their indices
	curr := original
	idx := 0
	for curr != nil {
		origMap[curr] = idx
		curr = curr.Next
		idx++
	}

	// Map copied nodes to their indices
	curr = copied
	idx = 0
	for curr != nil {
		copyMap[curr] = idx
		curr = curr.Next
		idx++
	}

	assert.Equal(t, len(origMap), len(copyMap), "Lists should have the exact same length")

	currOrig := original
	currCopy := copied

	// Traverse side by side to validate deep copy constraints
	for currOrig != nil {
		// 1. Must be different memory addresses (deep copy)
		assert.NotSame(t, currOrig, currCopy, "Nodes must be deeply copied (different memory addresses)")

		// 2. Values must match
		assert.Equal(t, currOrig.Val, currCopy.Val, "Node values must match")

		// 3. Validate Random pointers
		if currOrig.Random != nil {
			assert.NotNil(t, currCopy.Random, "Copied random pointer should not be nil")
			assert.NotSame(t, currOrig.Random, currCopy.Random, "Random pointers must point to newly created nodes")

			// Check if the random pointers resolve to the equivalent index in both lists
			origRandomIdx := origMap[currOrig.Random]
			copyRandomIdx := copyMap[currCopy.Random]
			message := fmt.Sprintf("Random pointer of node with value %d points to %d, but should point to %d",
				currCopy.Random.Val, copyRandomIdx, origRandomIdx)

			assert.Equal(t, origRandomIdx, copyRandomIdx, message)
		} else {
			assert.Nil(t, currCopy.Random, "Copied random pointer should be nil")
		}

		currOrig = currOrig.Next
		currCopy = currCopy.Next
	}
}
