package main

func main() {
}

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	prefix := prefixProduct(nums)
	sufix := sufixProduct(nums)

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result[i] = sufix[i+1]
			continue
		}

		if i == len(nums)-1 {
			result[i] = prefix[i-1]
			continue
		}

		result[i] = prefix[i-1] * sufix[i+1]
	}

	return result
}

func prefixProduct(nums []int) []int {
	result := make([]int, len(nums))
	result[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i]
	}

	return result
}

func sufixProduct(nums []int) []int {
	result := make([]int, len(nums))
	result[len(nums)-1] = nums[len(nums)-1]

	for i := len(nums) - 2; i >= 0; i-- {
		result[i] = result[i+1] * nums[i]
	}

	return result
}
