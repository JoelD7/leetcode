package container_most_water

import "math"

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	max := 0
	minHeight := 0
	area := 0

	for l < r {
		minHeight = height[l]
		if height[r] < height[l] {
			minHeight = height[r]
		}

		area = int(math.Abs(float64(l-r))) * minHeight
		if area > max {
			max = area
		}

		if minHeight == height[l] {
			l++
		} else {
			r--
		}
	}

	return max
}
