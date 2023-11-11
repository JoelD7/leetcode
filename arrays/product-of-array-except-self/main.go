package product_of_array_except_self

// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]
func productExceptSelf(nums []int) []int {
	result := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		val := 1
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}

			val = val * nums[j]
		}

		result = append(result, val)
	}

	return result
}
