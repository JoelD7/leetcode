package daily_temperatures

const initFlag = -1

// [73,74,75,71,69,72,76,73]
func dailyTemperatures(temperatures []int) []int {
	result := initResult(len(temperatures))

	stack := newStack()

	var cur, next, top, minDisFromCur, newTop, j int
	for i := 0; i < len(temperatures)-1; i++ {
		//Element has already been processed
		if result[i] != initFlag {
			continue
		}

		cur = temperatures[i]
		next = temperatures[i+1]

		if next > cur {
			result[i] = 1
			continue
		}

		if stack.length > 0 {
			minDisFromCur = 0

			for {
				top = stack.pop()
				newTop = stack.peek()

				if top > newTop {
					result[i+stack.length-1] = 1
				}

				if top > cur {
					minDisFromCur = 1
				} else {
					minDisFromCur++
				}

				//Only two elements in the stack: cur and next(top)
				if stack.length == 2 && top > cur {
					result[i] = 1
					break
				}

				if stack.length == 0 && minDisFromCur != 1 {
					result[i] = minDisFromCur
					break
				}
			}

			result[i] = minDisFromCur
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

func initResult(length int) []int {
	result := make([]int, length)

	for i := 0; i < length; i++ {
		result[i] = initFlag
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
