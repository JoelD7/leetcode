package daily_temperatures

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := newStack()

	for i := 0; i < len(temperatures)-1; i++ {
		if temperatures[i] < temperatures[i+1] {
			result[i] = 1

			for !stack.isEmpty() && temperatures[stack.peek()] < temperatures[i+1] {
				result[stack.peek()] = (i + 1) - stack.pop()
			}

		} else {
			stack.push(i)
		}
	}

	for !stack.isEmpty() {
		result[stack.pop()] = 0
	}

	return result
}

type node struct {
	val  int
	prev *node
}

type Stack struct {
	head   *node
	length int
}

func newStack() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) push(val int) {
	s.length++

	if s.head == nil || s.length == 0 {
		s.head = &node{val, nil}
		return
	}

	curHead := &node{s.head.val, s.head.prev}
	s.head.val = val
	s.head.prev = curHead
}

func (s *Stack) pop() int {
	s.length--

	curHead := &node{s.head.val, s.head.prev}
	s.head = curHead.prev

	return curHead.val
}

func (s *Stack) peek() int {
	return s.head.val
}

func (s *Stack) isEmpty() bool {
	return s.length == 0
}
