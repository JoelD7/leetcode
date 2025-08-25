package linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	visited := map[ListNode]struct{}{}
	var ok bool

	for head.Next != nil {
		_, ok = visited[*head]
		if ok {
			return true
		}

		visited[*head] = struct{}{}
		head = head.Next
	}

	return false
}
