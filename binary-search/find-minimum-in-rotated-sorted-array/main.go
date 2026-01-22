package find_minimum_in_rotated_sorted_array

func findMin(nums []int) int {
	left, mid, right := 0, 0, len(nums)-1

	for left < right {
		mid = left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
