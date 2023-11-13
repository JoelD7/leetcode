package product_of_array_except_self

import "fmt"

// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]
func productExceptSelf(nums []int) []int {
	result := make([]int, 0)

	fmt.Println(prefixProduct(nums))

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

// input: [1,2,3,4]
// output: [1,2,6,24]
func prefixProduct(nums []int) []int {
	prev := nums[0]
	result := make([]int, 0, len(nums))
	result = append(result, prev)

	for i := 1; i < len(nums); i++ {
		result = append(result, prev*nums[i])
		prev = prev * nums[i]
	}

	return result
}

// input: [1,2,3,4]
// output: [24,24,12,4]
func suffixProduct(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	next := nums[len(nums)-1]
	result := make([]int, len(nums))
	result[len(nums)-1] = next

	for i := len(nums) - 2; i >= 0; i-- {
		result[i] = next * nums[i]
		next = next * nums[i]
	}

	return result
}
