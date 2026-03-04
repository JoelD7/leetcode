package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var take, sum int
	resHead := &ListNode{0, nil}
	curRes := resHead

	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val + take

		if sum >= 10 {
			take = sum / 10
			curRes.Val = sum % 10
		} else {
			curRes.Val = sum
			take = 0
		}

		if l1.Next != nil || l2.Next != nil {
			curRes.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
			curRes = curRes.Next
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	rem := l1
	if rem == nil {
		rem = l2
	}

	for rem != nil {
		sum = rem.Val + take

		if sum >= 10 {
			take = sum / 10
			curRes.Val = sum % 10
		} else {
			curRes.Val = sum
			take = 0
		}

		if rem.Next != nil {
			curRes.Next = &ListNode{
				Val:  0,
				Next: nil,
			}
			curRes = curRes.Next
		}

		rem = rem.Next
	}

	if take > 0 {
		curRes.Next = &ListNode{
			Val:  0,
			Next: nil,
		}
		curRes = curRes.Next
		curRes.Val = take
	}

	return resHead
}
