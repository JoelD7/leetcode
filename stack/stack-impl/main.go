package main

type Stack struct {
	size int
	head *Node
}

type Node struct {
	val  int
	next *Node
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val int) {
	s.size++

	if s.head == nil {
		s.head = &Node{
			val: val,
		}
		return
	}

	oldHead := &Node{
		val:  s.head.val,
		next: s.head.next,
	}

	s.head.val = val
	s.head.next = oldHead
}

func (s *Stack) Pop() int {
	if s.size == 0 {
		return 0
	}

	s.size--

	val := s.head.val
	s.head = s.head.next
	return val
}

func (s *Stack) Peek() int {
	if s.size == 0 {
		return 0
	}

	return s.head.val
}
