package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Run("l1 = [5], l2 = [5]", func(t *testing.T) {
		l1 := BuildList(5)
		l2 := BuildList(5)

		assert.Equal(t, []int{0, 1}, listToArray(addTwoNumbers(l1, l2)))
	})
}

func BuildList(vals ...int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := NewListNode(vals[0])
	current := head

	for i := 1; i < len(vals); i++ {
		current.Next = NewListNode(vals[i])
		current = current.Next
	}

	return head
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}

func listToArray(list *ListNode) []int {
	head := list
	arr := make([]int, 0)

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	return arr
}
