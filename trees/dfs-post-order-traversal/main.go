package main

import (
	"github.com/JoelD7/leetcode/trees/utils"
)

func postOrderTraversalRecursive(node *utils.BinaryNode) []int {
	var result []int
	var recurse func(node *utils.BinaryNode, result *[]int)

	recurse = func(node *utils.BinaryNode, result *[]int) {
		if node == nil {
			return
		}

		recurse(node.Left, result)
		recurse(node.Right, result)

		*result = append(*result, node.Val)
	}

	recurse(node, &result)

	return result
}

type Stack struct {
	size int
	head *Node
}

type Node struct {
	val  *utils.BinaryNode
	next *Node
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val *utils.BinaryNode) {
	s.size++

	node := &Node{
		val: val,
	}

	if s.head == nil {
		s.head = node
	} else {
		curHead := s.head
		s.head = node
		s.head.next = curHead
	}
}

func (s *Stack) Pop() *utils.BinaryNode {
	s.size--
	if s.head == nil {
		return nil
	}

	val := s.head.val
	s.head = s.head.next
	return val
}

func postOrderTraversalIterative(node *utils.BinaryNode) []int {
	var result []int
	var cur *utils.BinaryNode

	stack1 := NewStack()
	stack1.Push(node)

	stack2 := NewStack()

	for stack1.size > 0 {
		cur = stack1.Pop()
		stack2.Push(cur)

		if cur.Left != nil {
			stack1.Push(cur.Left)
		}

		if cur.Right != nil {
			stack1.Push(cur.Right)
		}
	}

	for stack2.size > 0 {
		result = append(result, stack2.Pop().Val)
	}

	return result
}
