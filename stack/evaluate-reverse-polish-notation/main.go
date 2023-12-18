package evaluate_reverse_polish_notation

import (
	"errors"
	"strconv"
)

/*
Algorithm
- Push each token to stack
- When the token is an operator, pop the last two numbers and apply the operation to them
- Push the result of the operation to the stack
- Repeat
*/
func evalRPN(tokens []string) int {
	stack := NewStack()

	operators := map[string]struct{}{
		"+": {},
		"-": {},
		"*": {},
		"/": {},
	}

	var val, a, b int
	var err error
	var ok bool

	for _, token := range tokens {
		_, ok = operators[token]
		if ok {
			b = stack.Pop()
			a = stack.Pop()
			stack.Push(evaluate(a, b, token))
			continue
		}

		val, err = strconv.Atoi(token)
		if err != nil {
			panic(err)
		}

		stack.Push(val)
	}

	return stack.Pop()
}

func evaluate(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return -1
	}
}

type Node struct {
	val  int
	prev *Node
}

type Stack struct {
	head   *Node
	length int
}

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Push(val int) {
	s.length++

	if s.head == nil || s.length == 0 {
		s.head = &Node{val, nil}
		return
	}

	curHead := &Node{s.head.val, s.head.prev}
	s.head = &Node{val, curHead}
}

func (s *Stack) Pop() int {
	if s.head == nil || s.length == 0 {
		panic(errors.New("empty stack"))
	}

	s.length--

	curHead := &Node{s.head.val, s.head.prev}
	s.head = curHead.prev

	return curHead.val
}
