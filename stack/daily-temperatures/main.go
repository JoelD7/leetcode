package daily_temperatures

func dailyTemperatures(temperatures []int) []int {
	answer := make([]int, len(temperatures))
	stack := make([]int, 0)

	for i, temp := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temp {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[index] = i - index
		}

		stack = append(stack, i)
	}

	return answer
}
