package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeque(t *testing.T) {

	t.Run("New Deque is empty", func(t *testing.T) {
		d := NewDeque()
		assert.True(t, d.IsEmpty(), "New deque should be empty")
		assert.Equal(t, 0, d.size, "New deque size should be 0")
	})

	// 2. Test PushBack and PopBack (Stack Behavior at Tail)
	t.Run("PushBack and PopBack", func(t *testing.T) {
		d := NewDeque()
		d.PushBack(10)
		d.PushBack(20)

		assert.False(t, d.IsEmpty(), "Deque should not be empty")
		assert.Equal(t, 2, d.size, "Size should be 2")
		assert.Equal(t, 20, d.Back(), "Back element should be 20")

		assert.Equal(t, 20, d.PopBack(), "First PopBack should return 20")
		assert.Equal(t, 10, d.PopBack(), "Second PopBack should return 10")

		assert.True(t, d.IsEmpty(), "Deque should be empty after popping all elements")
	})

	t.Run("PushFront and PopFront", func(t *testing.T) {
		d := NewDeque()
		d.PushFront(10)
		d.PushFront(20)

		assert.Equal(t, 2, d.size, "Size should be 2")
		assert.Equal(t, 20, d.Front(), "Front element should be 20")

		assert.Equal(t, 20, d.PopFront(), "First PopFront should return 20")
		assert.Equal(t, 10, d.PopFront(), "Second PopFront should return 10")

		assert.True(t, d.IsEmpty(), "Deque should be empty")
	})

	t.Run("Queue Behavior (PushBack -> PopFront)", func(t *testing.T) {
		d := NewDeque()
		d.PushBack(1)
		d.PushBack(2)
		d.PushBack(3)

		assert.Equal(t, 1, d.PopFront(), "Should pop 1")
		assert.Equal(t, 2, d.PopFront(), "Should pop 2")
		assert.Equal(t, 3, d.PopFront(), "Should pop 3")
	})

	t.Run("Reverse Queue (PushFront -> PopBack)", func(t *testing.T) {
		d := NewDeque()
		d.PushFront(1)
		d.PushFront(2)
		d.PushFront(3)

		assert.Equal(t, 1, d.PopBack(), "Should pop 1 (first inserted)")
		assert.Equal(t, 2, d.PopBack(), "Should pop 2")
		assert.Equal(t, 3, d.PopBack(), "Should pop 3")
	})

	t.Run("Pop on Empty", func(t *testing.T) {
		d := NewDeque()
		assert.Equal(t, -1, d.PopFront(), "PopFront on empty should return -1")
		assert.Equal(t, -1, d.PopBack(), "PopBack on empty should return -1")
	})

	t.Run("Mixed Push/Pop", func(t *testing.T) {
		d := NewDeque()
		d.PushBack(1)  // [1]
		d.PushFront(2) // [2, 1]
		d.PushBack(3)  // [2, 1, 3]

		assert.Equal(t, 2, d.Front(), "Front should be 2")
		assert.Equal(t, 3, d.Back(), "Back should be 3")

		d.PopFront() // Removes 2 -> [1, 3]
		assert.Equal(t, 1, d.Front(), "Front after pop should be 1")

		d.PopBack() // Removes 3 -> [1]
		assert.Equal(t, 1, d.Back(), "Back after pop should be 1")
	})
}
