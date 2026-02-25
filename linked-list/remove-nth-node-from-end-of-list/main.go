package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	left, right := head, head
	var prevLeft *ListNode

	for i := 0; i < n; i++ {
		right = right.Next
	}

	if right == nil {
		left = left.Next
		return left
	}

	for right != nil {
		prevLeft = left
		left = left.Next
		right = right.Next
	}

	prevLeft.Next = left.Next

	return head
}
