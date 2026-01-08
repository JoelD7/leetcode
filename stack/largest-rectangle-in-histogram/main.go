package largest_rectangle_in_histogram

type Node struct {
	val  int
	next *Node
}

type Stack struct {
	size int
	head Node
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(val int) {
	s.size++
	prevHead := s.head
	s.head = Node{
		val:  val,
		next: &prevHead,
	}
}

func (s *Stack) Pop() int {
	if s.size == 0 {
		return 0
	}

	s.size--
	var newHead Node
	var val int
	if s.head.next != nil {
		newHead = *s.head.next
		val = s.head.val
	}

	s.head = newHead
	return val
}

func (s *Stack) Peek() int {
	return s.head.val
}

func largestRectangleArea(heights []int) int {
	stack := NewStack()
	stack.Push(-1)

	var maxArea int

	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	for i, _ := range heights {
		for stack.Peek() != -1 && heights[i] <= heights[stack.Peek()] {
			height := heights[stack.Pop()]
			width := i - stack.Peek() - 1
			maxArea = max(maxArea, height*width)
		}

		stack.Push(i)
	}

	for stack.Peek() != -1 {
		height := heights[stack.Pop()]
		width := len(heights) - stack.Peek() - 1
		maxArea = max(maxArea, height*width)
	}

	return maxArea
}
