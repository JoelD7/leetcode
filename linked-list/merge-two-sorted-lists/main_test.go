package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	t.Run("list1 = [1,2,4], list2 = [1,3,4]", func(t *testing.T) {
		list1 := BuildList(1, 2, 4)
		list2 := BuildList(1, 3, 4)

		output := mergeTwoLists(list1, list2)
		assert.Equal(t, []int{1, 1, 2, 3, 4, 4}, listToArray(output))
	})

	t.Run("list1 = [5], list2 = [1,2,4]", func(t *testing.T) {
		list1 := BuildList(5)
		list2 := BuildList(1, 2, 4)

		output := mergeTwoLists(list1, list2)
		assert.Equal(t, []int{1, 2, 4, 5}, listToArray(output))
	})
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
