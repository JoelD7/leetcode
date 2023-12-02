package container_most_water

import "math"

/*
Notes

- From the given example, the are appears to be wxh, where w = the numbers between the points + 1, and
h = height of the smaller of the two points.
- Cannot sort the array because order matters.
*/

func maxArea(height []int) int {
	//1. Find the largest number in height
	//2. Set one of the pointers to that index
	//3. Move the other pointer along height
	//4. Calculate the area on each iteration and save the max area in a variable, update it if you find another larger area
	largest := 0
	a := 0

	for i := 0; i < len(height); i++ {
		if height[i] > largest {
			largest = height[i]
			a = i
		}
	}

	max := 0
	area := 0
	minHeight := 0

	for i := 0; i < len(height); i++ {
		if i == a {
			continue
		}

		minHeight = height[i]
		if height[a] < height[i] {
			minHeight = height[a]
		}

		area = int(math.Abs(float64(i-a))) * minHeight
		if area > max {
			max = area
		}
	}

	return max
}
