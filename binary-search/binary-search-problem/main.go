package binary_search_problem

func search(nums []int, target int) int {
	min := 0
	max := len(nums) - 1

	for min <= max {
		middle := min + (max-min)/2

		if nums[middle] == target {
			return middle
		}

		if nums[middle] > target {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}

	return -1
}
