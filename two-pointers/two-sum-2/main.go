package two_sum_2

func twoSum(numbers []int, target int) []int {
	if len(numbers) == 0 {
		return []int{}
	}

	l := 0
	r := len(numbers) - 1

	for l < r {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		} else if numbers[l]+numbers[r] < target {
			l++
		} else {
			r--
		}
	}

	return []int{}
}
