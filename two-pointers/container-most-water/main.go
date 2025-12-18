package container_most_water

func maxArea(height []int) int {
	max := 0
	getMax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	getMin := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	area, l := 0, 0
	r := len(height) - 1

	for l < r {
		area = (r - l) * getMin(height[l], height[r])
		max = getMax(area, max)

		if height[r] < height[l] {
			r--
		} else {
			l++
		}
	}

	return max
}
