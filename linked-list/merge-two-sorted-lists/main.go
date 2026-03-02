package main

type ListNode struct {
	Val  int
	Next *ListNode
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

func SetNext(val int, l *ListNode) *ListNode {
	next := &ListNode{
		Val: val,
	}
	l.Next = next
	return next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var head *ListNode
	if list1.Val < list2.Val {
		head = list1
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}

	cur := head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	if list1 == nil {
		cur.Next = list2
	} else {
		cur.Next = list1
	}

	return head
}
