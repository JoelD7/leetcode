package valid_parentheses

type Node struct {
	val  byte
	next *Node
}

type Stack struct {
	size int
	head Node
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(val byte) {
	s.size++
	prevHead := s.head
	s.head = Node{
		val:  val,
		next: &prevHead,
	}
}

func (s *Stack) Pop() byte {
	if s.size == 0 {
		return 0
	}

	s.size--
	var newHead Node
	var val byte
	if s.head.next != nil {
		newHead = *s.head.next
		val = s.head.val
	}

	s.head = newHead
	return val
}

func (s *Stack) Peek() byte {
	return s.head.val
}

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}

	openToClosed := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	closedToOpen := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := NewStack()

	var c byte
	for i := 0; i < len(s); i++ {
		c = s[i]

		_, isOpen := openToClosed[c]
		if isOpen {
			stack.Push(c)
			continue
		}

		match, isClosed := closedToOpen[c]
		if isClosed && match != stack.Peek() {
			return false
		}

		if isClosed && match == stack.Peek() {
			stack.Pop()
		}

	}

	return stack.size == 0
}
