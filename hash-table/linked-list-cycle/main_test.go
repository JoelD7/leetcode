package linked_list_cycle

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHasCycle(t *testing.T) {
	c := require.New(t)

	t.Run("[3,2,0,-4]", func(t *testing.T) {
		node3 := &ListNode{Val: 3}
		node2 := &ListNode{Val: 2}
		node0 := &ListNode{Val: 0}
		node4 := &ListNode{Val: -4}
		node3.Next = node2
		node2.Next = node0
		node0.Next = node4
		node4.Next = node2

		c.True(hasCycle(node3))
	})

	t.Run("[1,2]", func(t *testing.T) {
		node1 := &ListNode{Val: 1}
		node2 := &ListNode{Val: 2}
		node1.Next = node2
		node2.Next = node1
		c.True(hasCycle(node1))
	})

	t.Run("[1]", func(t *testing.T) {
		node1 := &ListNode{Val: 1}
		c.False(hasCycle(node1))
	})

	t.Run("[-21,10,17,8,4,26,5,35,33,-7,-16,27,-12,6,29,-12,5,9,20,14,14,2,13,-24,21,23,-21,5]", func(t *testing.T) {
		node1 := &ListNode{Val: -21}
		node2 := &ListNode{Val: 10}
		node3 := &ListNode{Val: 17}
		node4 := &ListNode{Val: 8}
		node5 := &ListNode{Val: 4}
		node6 := &ListNode{Val: 26}
		node7 := &ListNode{Val: 5}
		node8 := &ListNode{Val: 35}
		node9 := &ListNode{Val: 33}
		node10 := &ListNode{Val: -7}
		node11 := &ListNode{Val: -16}
		node12 := &ListNode{Val: 27}
		node13 := &ListNode{Val: -12}
		node14 := &ListNode{Val: 6}
		node15 := &ListNode{Val: 29}
		node16 := &ListNode{Val: -12}
		node17 := &ListNode{Val: 5}
		node18 := &ListNode{Val: 9}
		node19 := &ListNode{Val: 20}
		node20 := &ListNode{Val: 14}
		node21 := &ListNode{Val: 14}
		node22 := &ListNode{Val: 2}
		node23 := &ListNode{Val: 13}
		node24 := &ListNode{Val: -24}
		node25 := &ListNode{Val: 21}
		node26 := &ListNode{Val: 23}
		node27 := &ListNode{Val: -21}
		node28 := &ListNode{Val: 5}
		node1.Next = node2
		node2.Next = node3
		node3.Next = node4
		node4.Next = node5
		node5.Next = node6
		node6.Next = node7
		node7.Next = node8
		node8.Next = node9
		node9.Next = node10
		node10.Next = node11
		node11.Next = node12
		node12.Next = node13
		node13.Next = node14
		node14.Next = node15
		node15.Next = node16
		node16.Next = node17
		node17.Next = node18
		node18.Next = node19
		node19.Next = node20
		node20.Next = node21
		node21.Next = node22
		node22.Next = node23
		node23.Next = node24
		node24.Next = node25
		node25.Next = node26
		node26.Next = node27
		node27.Next = node28
		c.False(hasCycle(node1))
	})
}
