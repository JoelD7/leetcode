package product_of_array_except_self

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	prefix := prefixProduct(nums)
	suffix := suffixProduct(nums)

	result[0] = suffix[1]
	result[len(nums)-1] = prefix[len(nums)-2]

	for i := 1; i < len(nums)-1; i++ {
		result[i] = prefix[i-1] * suffix[i+1]
	}

	return result
}

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
