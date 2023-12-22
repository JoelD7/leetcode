package remove_outermost_parentheses

import "strings"

func removeOuterParentheses(s string) string {
	primitives := make([]string, 0)

	var str, tmp string
	stack := NewStack()

	for i, c := range s {
		str = string(c)
		tmp += str

		if str == "(" {
			stack.Push(0)
		} else {
			stack.Pop()
		}

		if stack.length == 0 && i != 0 {
			tmp = tmp[1 : len(tmp)-1]
			primitives = append(primitives, tmp)
			tmp = ""
		}
	}

	return strings.Join(primitives, "")
}

type Node struct {
	val  int
	prev *Node
}

type Stack struct {
	length int
	head   *Node
}

func NewStack() *Stack {
	return &Stack{0, nil}
}

func (s *Stack) Push(val int) {
	s.length++

	if s.head == nil {
		s.head = &Node{val, nil}
		return
	}

	curHead := &Node{s.head.val, s.head.prev}
	s.head = &Node{val, curHead}
}

func (s *Stack) Pop() int {
	s.length--

	curHead := &Node{s.head.val, s.head.prev}
	s.head = curHead.prev

	return curHead.val
}
