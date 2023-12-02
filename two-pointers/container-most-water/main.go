package container_most_water

import "math"

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

	max := 0
	minHeight := 0
	area := 0

	for i := 0; i < len(height); i++ {
		for j := 0; j < len(height); j++ {
			if i == j {
				continue
			}

			minHeight = height[i]
			if height[j] < height[i] {
				minHeight = height[j]
			}

			area = int(math.Abs(float64(i-j))) * minHeight
			if area > max {
				max = area
			}
		}
	}

	return max
}
