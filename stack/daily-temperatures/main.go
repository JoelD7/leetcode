package daily_temperatures

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))

	for i, temp := range temperatures {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temp {
				result[i] = j - i
				break
			}
		}
	}

	return result
}
