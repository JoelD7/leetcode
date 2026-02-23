package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	l := head
	r := head.Next.Next
	for r != nil && r.Next != nil {
		if l == r {
			return true
		}

		r = r.Next.Next
		l = l.Next
	}

	return false
}
