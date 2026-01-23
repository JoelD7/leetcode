package search_rotated_sorted_array

func search(nums []int, target int) int {
	left, mid, right := 0, 0, len(nums)-1

	for left < right {
		mid = left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	rot := left
	left, mid, right = 0, 0, len(nums)-1

	if target >= nums[rot] && target <= nums[right] {
		left = rot
	} else {
		right = rot
	}

	for left <= right {
		mid = left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
