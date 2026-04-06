package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	size int
	head *Node
	tail *Node
}

type Node struct {
	val  *TreeNode
	next *Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) enqueue(val *TreeNode) {
	node := &Node{val: val}

	if q.size == 0 {
		q.head = node
		q.tail = node
		q.size++
		return
	}

	curTail := q.tail
	q.tail = node
	curTail.next = node

	q.size++
}

func (q *Queue) dequeue() *TreeNode {
	val := q.head.val
	q.head = q.head.next
	q.size--
	return val
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	if root == nil {
		return res
	}

	queue := NewQueue()
	queue.enqueue(root)

	for queue.size > 0 {
		nodesPerLvl := make([]int, queue.size)
		var node *TreeNode

		for i := 0; i < len(nodesPerLvl); i++ {
			node = queue.dequeue()
			nodesPerLvl[i] = node.Val

			if node.Left != nil {
				queue.enqueue(node.Left)
			}

			if node.Right != nil {
				queue.enqueue(node.Right)
			}
		}

		res = append(res, nodesPerLvl)
	}

	return res
}
