package main

import (
	"github.com/JoelD7/leetcode/trees/utils"
)

func bfsTraversalRecursive(root *utils.TreeNode) [][]int {
	var nodesByLevel [][]int
	var recurse func(node *utils.TreeNode, level int, nodesByLevel *[][]int)

	recurse = func(node *utils.TreeNode, level int, nodesByLevel *[][]int) {
		if node == nil {
			return
		}

		if len(*nodesByLevel) <= level {
			*nodesByLevel = append(*nodesByLevel, []int{})
		}

		(*nodesByLevel)[level] = append((*nodesByLevel)[level], node.Val)
		recurse(node.Left, level+1, nodesByLevel)
		recurse(node.Right, level+1, nodesByLevel)
	}

	recurse(root, 0, &nodesByLevel)

	return nodesByLevel
}

type Queue struct {
	size int
	head *Node
	tail *Node
}

type Node struct {
	val  *utils.TreeNode
	next *Node
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) enqueue(val *utils.TreeNode) {
	q.size++

	node := &Node{val: val, next: nil}

	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		prevTail := q.tail
		q.tail = node
		prevTail.next = q.tail
	}
}

func (q *Queue) dequeue() *utils.TreeNode {
	q.size--

	if q.head == nil {
		return nil
	}

	val := q.head.val
	q.head = q.head.next

	return val
}

func bfsTraversalIterative(root *utils.TreeNode) [][]int {
	var result [][]int
	var levelWidth int
	var cur *utils.TreeNode

	queue := NewQueue()
	queue.enqueue(root)

	for queue.size > 0 {
		nodesInLevel := make([]int, 0)
		levelWidth = queue.size

		for i := 0; i < levelWidth && queue.size > 0; i++ {
			cur = queue.dequeue()
			nodesInLevel = append(nodesInLevel, cur.Val)

			if cur.Left != nil {
				queue.enqueue(cur.Left)
			}

			if cur.Right != nil {
				queue.enqueue(cur.Right)
			}
		}

		result = append(result, nodesInLevel)
	}

	return result
}
