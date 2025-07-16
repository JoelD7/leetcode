package search_rotated_sorted_array

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	mid := 0

	for low <= high {
		mid = (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target && nums[low] > target && nums[low] > nums[mid] {
			high = mid - 1
		} else if nums[mid] > target && nums[low] > target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target && nums[low] > nums[mid] && nums[low] > target {
			low = mid + 1
		} else if nums[mid] < target && nums[low] > nums[mid] {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		}
	}

	return -1
}
