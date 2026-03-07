package main

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	arr := make([]int, 0)

	var cur *ListNode
	for _, list := range lists {
		cur = list
		for cur != nil {
			arr = append(arr, cur.Val)
			cur = cur.Next
		}
	}

	sort.Ints(arr)

	var result *ListNode
	cur = result
	for i, item := range arr {
		if cur != nil {
			cur.Val = item
		} else {
			result = &ListNode{
				Val: item,
			}
			cur = result
		}

		if i < len(arr)-1 {
			cur.Next = &ListNode{}
		}
		cur = cur.Next
	}

	return result
}
