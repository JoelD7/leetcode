package evaluate_reverse_polish_notation

import (
	"strconv"
)

type Stack struct {
	size int
	head Node
}

type Node struct {
	val  int
	next *Node
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(val int) {
	if s.size == 0 {
		s.size++
		s.head = Node{
			val: val,
		}
	}

	s.size++

	curHead := s.head
	s.head = Node{
		val:  val,
		next: &curHead,
	}
}

func (s *Stack) Pop() int {
	s.size--
	val := s.head.val

	var newHead Node
	if s.head.next != nil {
		newHead = *s.head.next
	}

	s.head = newHead
	return val
}

func evalRPN(tokens []string) int {
	var token string
	var opOne, opTwo, tokenInt int
	operators := map[string]struct{}{
		"+": {},
		"-": {},
		"*": {},
		"/": {},
	}

	stack := NewStack()

	for i := 0; i < len(tokens); i++ {
		token = tokens[i]

		_, isOperator := operators[token]
		if !isOperator {
			//Ignore error because input always will have valid integers
			tokenInt, _ = strconv.Atoi(token)
			stack.Push(tokenInt)
			continue
		}

		opTwo = stack.Pop()
		opOne = stack.Pop()

		switch token {
		case "+":
			stack.Push(opOne + opTwo)
		case "-":
			stack.Push(opOne - opTwo)
		case "*":
			stack.Push(opOne * opTwo)
		case "/":
			stack.Push(opOne / opTwo)
		}
	}

	return stack.Pop()
}
