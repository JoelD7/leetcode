package container_most_water

func maxArea(height []int) int {
	max := 0
	getMax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var w, h int
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			w = j - i
			h = height[i]
			if height[j] < height[i] {
				h = height[j]
			}

			max = getMax(max, w*h)
		}
	}

	return max
}
