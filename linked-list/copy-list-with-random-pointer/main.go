package copy_list_with_random_pointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	og := head
	var next, copyNode *Node

	for og != nil {
		next = og.Next

		copyNode = &Node{
			Val:  og.Val,
			Next: next,
		}

		og.Next = copyNode
		og = next
	}

	cur := head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}

		cur = cur.Next.Next
	}

	og = head
	var nextOg *Node
	dummyHead := &Node{Val: 0}
	copyNode, copyOg := dummyHead, dummyHead

	for og != nil {
		nextOg = og.Next.Next

		copyNode = og.Next
		copyOg.Next = copyNode
		copyOg = copyNode

		og.Next = nextOg
		og = nextOg
	}

	return dummyHead.Next
}
