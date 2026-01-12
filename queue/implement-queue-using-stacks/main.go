package implement_queue_using_stacks

type Stack struct {
	head *Node
	size int
}

type Node struct {
	next *Node
	val  int
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(x int) {
	curHead := s.head

	if s.size == 0 {
		s.head = &Node{
			val: x,
		}
	} else {
		s.head = &Node{
			val:  x,
			next: curHead,
		}
	}

	s.size++
}

func (s *Stack) Pop() int {
	val := s.head.val
	s.head = s.head.next
	s.size--
	return val
}

func (s *Stack) Peek() int {
	return s.head.val
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

type MyQueue struct {
	s1   Stack
	size int
}

func Constructor() MyQueue {
	return MyQueue{
		s1:   NewStack(),
		size: 0,
	}
}

func (this *MyQueue) Push(x int) {
	s2 := NewStack()

	for !this.s1.IsEmpty() {
		s2.Push(this.s1.Pop())
	}

	s2.Push(x)

	for !s2.IsEmpty() {
		this.s1.Push(s2.Pop())
	}

	this.size++
}

func (this *MyQueue) Pop() int {
	this.size--
	return this.s1.Pop()
}

func (this *MyQueue) Peek() int {
	return this.s1.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.size == 0
}
