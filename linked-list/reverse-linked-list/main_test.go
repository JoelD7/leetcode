package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper function to easily construct a linked list for testing
func buildList(vals ...int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	for i := 1; i < len(vals); i++ {
		curr.Next = &ListNode{Val: vals[i]}
		curr = curr.Next
	}
	return head
}

func TestReverseList(t *testing.T) {
	t.Run("List with 5 elements", func(t *testing.T) {
		input := buildList(1, 2, 3, 4, 5)
		expected := buildList(5, 4, 3, 2, 1)

		assert.Equal(t, expected, reverseList(input))
	})

	t.Run("List with 2 elements", func(t *testing.T) {
		input := buildList(1, 2)
		expected := buildList(2, 1)

		assert.Equal(t, expected, reverseList(input))
	})

	t.Run("Empty list", func(t *testing.T) {
		var input *ListNode = nil
		var expected *ListNode = nil

		assert.Equal(t, expected, reverseList(input))
	})

	t.Run("List with 1 element", func(t *testing.T) {
		input := buildList(1)
		expected := buildList(1)

		assert.Equal(t, expected, reverseList(input))
	})

	t.Run("List with negative and positive elements", func(t *testing.T) {
		input := buildList(-3, 0, 9, -12, 5)
		expected := buildList(5, -12, 9, 0, -3)

		assert.Equal(t, expected, reverseList(input))
	})
}
