package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromTheEnd(t *testing.T) {
	t.Run("head = [1,2,3,4,5], n = 2", func(t *testing.T) {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		}

		result := removeNthFromEnd(head, 2)
		assert.Equal(t, []int{1, 2, 3, 5}, listToArray(result))
	})

	t.Run("head = [1], n = 1", func(t *testing.T) {
		head := &ListNode{
			Val: 1,
		}

		result := removeNthFromEnd(head, 1)
		assert.Equal(t, []int{}, listToArray(result))
	})

	t.Run("head = [1,2], n = 2", func(t *testing.T) {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		}

		result := removeNthFromEnd(head, 2)
		assert.Equal(t, []int{2}, listToArray(result))
	})

	t.Run("head = [1,2,3], n = 3", func(t *testing.T) {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
				},
			},
		}

		result := removeNthFromEnd(head, 3)
		assert.Equal(t, []int{2, 3}, listToArray(result))
	})

	t.Run("head = [1,2,3,4], n = 4", func(t *testing.T) {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
		}

		result := removeNthFromEnd(head, 4)
		assert.Equal(t, []int{2, 3, 4}, listToArray(result))
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
