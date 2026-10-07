package main

func reverseList(head *ListNode) *ListNode {
	var prev, cur, next *ListNode

	cur = head

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}
