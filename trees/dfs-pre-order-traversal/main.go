package main

import (
	"github.com/JoelD7/leetcode/trees/utils"
)

func preOrderTraversalRecursive(node *utils.BinaryNode) []int {
	var result []int

	var recurse func(node *utils.BinaryNode, result *[]int)

	recurse = func(node *utils.BinaryNode, result *[]int) {
		if node == nil {
			return
		}

		*result = append(*result, node.Val)
		recurse(node.Left, result)
		recurse(node.Right, result)
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

func preOrderTraversalIterative(node *utils.BinaryNode) []int {
	var result []int

	var cur *utils.BinaryNode

	stack := NewStack()
	stack.Push(node)

	for stack.size > 0 {
		cur = stack.Pop()
		result = append(result, cur.Val)

		if cur.Right != nil {
			stack.Push(cur.Right)
		}

		if cur.Left != nil {
			stack.Push(cur.Left)
		}
	}

	return result
}
