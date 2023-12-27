package daily_temperatures

// [73,74,75,71,69,72,76,73]
func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))

	stack := newStack()

	var cur, next, top, popCount, j int
	for i := 0; i < len(temperatures)-1; i++ {
		cur = temperatures[i]
		next = temperatures[i+1]

		if next > cur {
			result[i] = 1
			continue
		}

		if stack.length > 0 {
			for {
				top = stack.pop()
				popCount++

				if top > cur {

				}

				if top == cur {
					break
				}
			}

			result[i] = popCount
			continue
		}

		j = i + 1
		for {
			next = temperatures[j]
			if next > cur {
				result[i] = stack.length + 1
				break
			}

			stack.push(next)
			j++
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
