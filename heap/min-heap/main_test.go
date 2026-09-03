package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	h := &MinHeap{}

	// Case 1: Insert a single element
	h.Insert(10)
	assert.Equal(t, 10, h.Peek(), "Peek should return the only element (10)")

	// Case 2: Insert a smaller element (should become the new root)
	h.Insert(5)
	assert.Equal(t, 5, h.Peek(), "Peek should return the new minimum (5)")

	// Case 3: Insert multiple elements
	h.Insert(15)
	h.Insert(2)
	h.Insert(8)
	// Logical elements in heap: 2, 5, 8, 10, 15
	assert.Equal(t, 2, h.Peek(), "Peek should return the absolute minimum (2)")

	// Case 4: Sequential removals to ensure order is maintained
	assert.Equal(t, 2, h.Remove(), "First removal should return 2")
	assert.Equal(t, 5, h.Remove(), "Second removal should return 5")
	assert.Equal(t, 8, h.Remove(), "Third removal should return 8")

	// Case 5: Peek after multiple removals
	assert.Equal(t, 10, h.Peek(), "After removing 2, 5, and 8, the new minimum should be 10")

	// Case 6: Handle duplicate values
	h.Insert(10) // Insert a duplicate of the current minimum
	assert.Equal(t, 10, h.Peek(), "Peek should still return 10")
	assert.Equal(t, 10, h.Remove(), "Removal should return the first 10")
	assert.Equal(t, 10, h.Remove(), "Removal should return the second 10")

	// Case 7: Remove the last remaining element
	assert.Equal(t, 15, h.Remove(), "Removal should return the last element (15)")

	// Optional Case 8: Mixed operations
	h.Insert(100)
	h.Insert(20)
	assert.Equal(t, 20, h.Remove(), "Should remove 20")
	h.Insert(10)
	assert.Equal(t, 10, h.Remove(), "Should remove 10")
	assert.Equal(t, 100, h.Remove(), "Should remove 100")
}
