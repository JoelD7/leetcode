package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head.Next == nil {
		return head
	}

	var outHead, prevRevTail *ListNode
	cur := head
	secHead := head
	var count int

	for cur != nil {
		for count < k-1 {
			cur = cur.Next
			count++
		}

		if cur == nil {
			break
		}

		next := cur.Next
		cur.Next = nil

		revHead, revTail := reverse(secHead)
		if outHead == nil {
			outHead = revHead
		}

		if prevRevTail != nil {
			prevRevTail.Next = revHead
		}

		prevRevTail = revTail

		secHead = next
		cur = next
		count = 0
	}

	return nil
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}

	revTail := head
	cur := head

	var prev, next *ListNode
	for cur != nil {
		next = cur.Next

		cur.Next = prev
		prev = cur

		cur = next
	}

	return prev, revTail
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
