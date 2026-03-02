package utils

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
