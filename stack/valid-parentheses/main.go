package valid_parentheses

/*
- Itera sobre el string y ve agregando cada caracter c a un stack
- Si c es un opening tag agregalo al stack
- Si c es closing tag, haz pop del stack y ese elemento debe corresopnder con el closing tag, si no lo hace, return false.
- Al final de la iteracion return true.
*/
func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}

	openingTag := map[string]struct{}{
		"(": {},
		"{": {},
		"[": {},
	}

	closingTag := map[string]struct{}{
		")": {},
		"}": {},
		"]": {},
	}

	closedToOpen := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	var char, poped, openChar string
	var ok bool
	stack := newStack()

	for _, c := range s {
		char = string(c)

		_, ok = openingTag[char]
		if ok {
			stack.push(char)
			continue
		}

		_, ok = closingTag[char]
		if ok && stack.length == 0 {
			return false
		}

		if ok {
			poped = stack.pop()
		}

		openChar, _ = closedToOpen[char]
		if openChar != poped {
			return false
		}
	}

	if stack.length > 0 {
		return false
	}

	return true
}

type node struct {
	val  string
	prev *node
}

type Stack struct {
	head   *node
	length int
}

func newStack() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) push(val string) {
	s.length++

	if s.head == nil {
		s.head = &node{val, nil}

		return
	}

	curHead := &node{s.head.val, s.head.prev}
	s.head.val = val
	s.head.prev = curHead
}

func (s *Stack) pop() string {
	if s.head == nil {
		return ""
	}

	s.length--

	curHead := s.head
	s.head = curHead.prev

	return curHead.val
}

func (s *Stack) peek() string {
	if s.head == nil {
		return ""
	}

	return s.head.val
}
