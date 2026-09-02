package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap(t *testing.T) {
	// --- Scenario 1: Insert and Peek ---
	t.Run("Insert and Peek", func(t *testing.T) {
		h := &MaxHeap{}

		h.Insert(10)
		assert.Equal(t, 10, h.Peek(), "Peek should return the only inserted element")

		h.Insert(20)
		assert.Equal(t, 20, h.Peek(), "Peek should return the new max element")

		h.Insert(5)
		assert.Equal(t, 20, h.Peek(), "Peek should still return 20 after inserting a smaller element")

		h.Insert(-10)
		assert.Equal(t, 20, h.Peek(), "Peek should handle negative numbers correctly")
	})

	// --- Scenario 2: Remove (Descending Order) ---
	t.Run("Remove items in descending order", func(t *testing.T) {
		h := &MaxHeap{}
		elements := []int{3, 10, 5, 20, 1, 15}

		for _, el := range elements {
			h.Insert(el)
		}

		expected := []int{20, 15, 10, 5, 3, 1}
		for _, exp := range expected {
			assert.Equal(t, exp, h.Remove(), "Remove should extract elements in descending order")
		}
	})

	// --- Scenario 3: Handle Duplicates ---
	t.Run("Handle duplicates correctly", func(t *testing.T) {
		h := &MaxHeap{}
		elements := []int{5, 10, 5, 20, 10}

		for _, el := range elements {
			h.Insert(el)
		}

		expected := []int{20, 10, 10, 5, 5}
		for _, exp := range expected {
			assert.Equal(t, exp, h.Remove(), "Remove should extract duplicate elements properly")
		}
	})

	// --- Scenario 4: Empty Heap Panics ---
	t.Run("Panics on empty heap", func(t *testing.T) {
		h := &MaxHeap{}

		// Assuming your implementation panics on empty slices.
		// If it returns a default value/error instead, change this to assert.Equal or assert.Error.
		assert.Panics(t, func() { h.Peek() }, "Calling Peek() on an empty heap should panic")
		assert.Panics(t, func() { h.Remove() }, "Calling Remove() on an empty heap should panic")
	})
}
