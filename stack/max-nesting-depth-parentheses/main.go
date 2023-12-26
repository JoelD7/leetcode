package max_nesting_depth_parentheses

import "errors"

func maxDepth(s string) int {
	stack := NewStack()

	var char string
	var max int

	for _, c := range s {
		char = string(c)

		if char == "(" {
			stack.Push(char)
		} else if char == ")" {
			stack.Pop()
		}

		if stack.length > max {
			max = stack.length
		}
	}

	return max
}

type Node struct {
	val  string
	prev *Node
}

type Stack struct {
	length int
	head   *Node
}

func NewStack() *Stack {
	return &Stack{0, nil}
}

func (s *Stack) Push(val string) {
	s.length++

	if s.length == 0 || s.head == nil {
		s.head = &Node{val, nil}
		return
	}

	curHead := &Node{s.head.val, s.head.prev}
	s.head.val = val
	s.head.prev = curHead
}

func (s *Stack) Pop() string {
	if s.length == 0 || s.head == nil {
		panic(errors.New("empty stack"))
	}

	s.length--

	curHead := &Node{s.head.val, s.head.prev}
	s.head = curHead.prev

	return curHead.val
}
