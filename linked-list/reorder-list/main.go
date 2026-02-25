package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev, next *ListNode
	cur := slow.Next
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	slow.Next = nil

	originalHead, reversedHead := head, prev
	for reversedHead != nil {
		next = originalHead.Next
		originalHead.Next = reversedHead
		originalHead = reversedHead
		reversedHead = next
	}
}

func reverseList(head *ListNode) ListNode {
	var prev, next, cur *ListNode
	cur = head

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur

		cur = next
	}

	return *prev
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
