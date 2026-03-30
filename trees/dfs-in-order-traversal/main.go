package main

import (
	"github.com/JoelD7/leetcode/trees/utils"
)

func inOrderTraversalRecursive(root *utils.TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	recurse(root, &result)

	return result
}

func recurse(node *utils.TreeNode, result *[]int) {
	if node == nil {
		return
	}

	recurse(node.Left, result)

	*result = append(*result, node.Val)

	recurse(node.Right, result)
}

type Stack struct {
	size int
	head *Node
}

type Node struct {
	val  *utils.TreeNode
	next *Node
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val *utils.TreeNode) {
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

func (s *Stack) Pop() *utils.TreeNode {
	s.size--
	if s.head == nil {
		return nil
	}

	val := s.head.val
	s.head = s.head.next
	return val
}

func inOrderTraversalIterative(root *utils.TreeNode) []int {
	var result []int
	stack := NewStack()
	cur := root

	for cur != nil || stack.size > 0 {
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}

		cur = stack.Pop()
		result = append(result, cur.Val)

		cur = cur.Right
	}

	return result
}
