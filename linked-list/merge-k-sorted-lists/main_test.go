package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeKList(t *testing.T) {
	t.Run("lists = [[1,4,5],[1,3,4],[2,6]]", func(t *testing.T) {
		lists := []*ListNode{
			BuildList(1, 4, 5),
			BuildList(1, 3, 4),
			BuildList(2, 6),
		}

		assert.Equal(t, []int{1, 1, 2, 3, 4, 4, 5, 6}, listToArray(mergeKLists(lists)))
	})
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
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

func listToArray(list *ListNode) []int {
	head := list
	arr := make([]int, 0)

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	return arr
}
