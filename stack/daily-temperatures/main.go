package daily_temperatures

// [73,74,75,71,69,72,76,73]
func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))

	stack := newStack()

	var top, j int

	for i := 0; i < len(temperatures)-1; i++ {
		if !stack.isEmpty() {
			top = stack.peek()

			if temperatures[top] < temperatures[top+1] {
				result[stack.pop()] = 1
			}
		}

		j = i + 1
		for temperatures[j] <= temperatures[i] && j < len(temperatures)-1 {
			stack.push(j)
			j++
		}

		if temperatures[j] > temperatures[i] {
			result[i] = j - i
		}

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
