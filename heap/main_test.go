package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap(t *testing.T) {

	t.Run("Insert updates length and correct root", func(t *testing.T) {
		h := &MaxHeap{}
		h.Insert(10)
		h.Insert(20)
		h.Insert(5)
		assert.Equal(t, 5, h.arr[len(h.arr)-1])
		h.Insert(30)
		assert.Len(t, h.arr, 4)
		assert.Equal(t, 30, h.Peek())
	})

	t.Run("Verifies swapping", func(t *testing.T) {
		h := &MaxHeap{}
		h.Insert(43)
		h.Insert(35)
		h.Insert(20)
		assert.Equal(t, 43, h.Peek())
		assert.Equal(t, 35, h.arr[1]) //left child
		assert.Equal(t, 20, h.arr[2]) //right child

		h.Insert(71)
		h.Insert(19)
		assert.Len(t, h.arr, 5)
		assert.Equal(t, 71, h.Peek())            //new max
		assert.Equal(t, 43, h.arr[1])            //left child of new root, swap happens only in the left subtree
		assert.Equal(t, 20, h.arr[2])            //right child of new root, no position changed
		assert.Equal(t, 19, h.arr[len(h.arr)-1]) //Last added element is at last position

		h.Insert(50)
		assert.Len(t, h.arr, 6)
		//50 is placed at the last position, then it swaps places with 20
		assert.Equal(t, 50, h.arr[2])
		assert.Equal(t, 20, h.arr[len(h.arr)-1])
	})

	t.Run("Remove returns elements in descending order", func(t *testing.T) {
		h := &MaxHeap{}
		elements := []int{15, 3, 17, 20, 8, 25}
		for _, el := range elements {
			h.Insert(el)
		}

		expected := []int{25, 20, 17, 15, 8, 3}
		for _, exp := range expected {
			assert.Equal(t, exp, h.Remove())
		}
	})

	t.Run("Handle duplicates correctly", func(t *testing.T) {
		h := &MaxHeap{}
		h.Insert(10)
		h.Insert(10)
		h.Insert(5)
		assert.Equal(t, 10, h.Peek())
		val1 := h.Remove()
		val2 := h.Remove()
		assert.Equal(t, val1, val2)
	})

	t.Run("Insert and remove single element", func(t *testing.T) {
		h := &MaxHeap{}
		h.Insert(42)

		val := h.Remove()
		if val != 42 {
			t.Errorf("expected to remove 42, got %d", val)
		}
		if len(h.arr) != 0 {
			t.Errorf("expected length 0 after removal, got %d", len(h.arr))
		}
	})

	t.Run("Insert already sorted descending elements", func(t *testing.T) {
		h := &MaxHeap{}
		h.Insert(50)
		h.Insert(40)
		h.Insert(30)
		h.Insert(20)

		if h.arr[0] != 50 {
			t.Errorf("expected root to be 50, got %d", h.arr[0])
		}
	})

	t.Run("Insert already sorted ascending elements", func(t *testing.T) {
		h := &MaxHeap{}
		h.Insert(1)
		h.Insert(2)
		h.Insert(3)
		h.Insert(4)

		if h.arr[0] != 4 {
			t.Errorf("expected root to be 4, got %d", h.arr[0])
		}
	})
}
