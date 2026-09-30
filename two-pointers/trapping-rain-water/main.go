package main

func trap(height []int) int {
	var total int
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	l, r := 0, len(height)-1
	leftMax, rightMax := height[l], height[r]

	for l < r {
		if leftMax < rightMax {
			l++
			leftMax = max(leftMax, height[l])
			total += leftMax - height[l]
		} else {
			r--
			rightMax = max(rightMax, height[r])
			total += rightMax - height[r]
		}
	}

	return total
}
