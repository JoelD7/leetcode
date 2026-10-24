package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// makeList is a helper function to easily generate linked lists from slices for testing.
func makeList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	curr := head
	for i := 1; i < len(nums); i++ {
		curr.Next = &ListNode{Val: nums[i]}
		curr = curr.Next
	}
	return head
}

func TestAddTwoNumbers(t *testing.T) {
	t.Run("equal length lists without carry at the end", func(t *testing.T) {
		l1 := makeList([]int{2, 4, 3})
		l2 := makeList([]int{5, 6, 4})
		expected := makeList([]int{7, 0, 8})

		assert.Equal(t, expected, addTwoNumbers(l1, l2))
	})

	t.Run("single zero node", func(t *testing.T) {
		l1 := makeList([]int{0})
		l2 := makeList([]int{0})
		expected := makeList([]int{0})

		assert.Equal(t, expected, addTwoNumbers(l1, l2))
	})

	t.Run("different length lists with cascading carry", func(t *testing.T) {
		l1 := makeList([]int{9, 9, 9, 9, 9, 9, 9})
		l2 := makeList([]int{9, 9, 9, 9})
		expected := makeList([]int{8, 9, 9, 9, 0, 0, 0, 1})

		assert.Equal(t, expected, addTwoNumbers(l1, l2))
	})

	t.Run("lists resulting in an extra carry node at the end", func(t *testing.T) {
		l1 := makeList([]int{2, 4, 9})
		l2 := makeList([]int{5, 6, 4})
		expected := makeList([]int{7, 0, 4, 1})

		assert.Equal(t, expected, addTwoNumbers(l1, l2))
	})
}
