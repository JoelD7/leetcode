package k_radius_subarray_averages

func getAverages(nums []int, k int) []int {
	result := make([]int, len(nums))

	if k == 0 {
		return nums
	}

	for i := 0; i < len(nums); i++ {
		result[i] = -1
	}

	var j, total, prevStart int
	windowSize := (k * 2) + 1
	windowStart := 0

	for i := k; i < len(nums)-k; i++ {
		if i == k {
			for ; j < windowSize; j++ {
				total += nums[j]
			}

			prevStart = nums[windowStart]
			windowStart++
			result[i] = total / windowSize
			continue
		}

		total -= prevStart
		total += nums[i+k]
		prevStart = nums[windowStart]
		windowStart++
		result[i] = total / windowSize
	}

	return result
}
