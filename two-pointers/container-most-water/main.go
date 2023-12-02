package container_most_water

/*
Notes

- From the given example, the are appears to be wxh, where w = the numbers between the points + 1, and
h = height of the smaller of the two points.
- Cannot sort the array because order matters.
*/

func maxArea(height []int) int {
	//	Use two pointers. Advance the l pointer until there is no element between it and r.
	//	Then advance r until there is no element between it an l.
	//	Calculate the area on each iteration and save the max area in a variable, update it if you find another larger area

	l := 0
	r := len(height) - 1
	max := 0
	area := 0
	minHeight := 0

	for l < r {
		minHeight = height[r]
		if height[l] < height[r] {
			minHeight = height[l]
		}

		area = (r - l) * minHeight
		if area > max {
			max = area
		}

		l++
	}

	l = 0
	for r > l {
		minHeight = height[r]
		if height[l] < height[r] {
			minHeight = height[l]
		}

		area = (r - l) * minHeight
		if area > max {
			max = area
		}
		r--
	}

	return max
}
