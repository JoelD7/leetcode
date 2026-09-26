package main

func trap(height []int) int {
	var i, total int
	var min func(a, b int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	//Ignore these starting heights as they don't add to the result
	for height[i] == 0 {
		i++
	}

	j := i + 1

	for i < len(height) && j < len(height) {
		//heights between i and j that need to be subtracted when finding a valid container
		mids := make([]int, 0)

		//Move j until we find a bar "j" that can trap water with bar "i"
		for j < len(height) && height[i] > height[j] {
			mids = append(mids, height[j])
			j++
		}

		if j >= len(height) {
			break
		}

		//h has to be the min between the two to prevent overflowing of water
		h := min(height[i], height[j])
		w := j - i - 1
		raw := h * w
		for _, mid := range mids {
			raw -= mid
		}
		total += raw

		//reset
		i = j
		j = i + 1
	}

	return total
}
