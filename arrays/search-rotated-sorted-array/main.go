package search_rotated_sorted_array

// [5,1,2,3,4], target: 1
func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	mid := 0

	for low <= high {
		mid = (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target && nums[low] > target {
			if nums[low] > nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}

		} else if nums[mid] > target {
			high = mid - 1
			//This means that the pivot is on this range. Example: target: 7, nums: [6,7,0,1,2,4,5]
		} else if nums[mid] < target && nums[low] > nums[mid] {
			if nums[low] > target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else if nums[mid] < target {
			low = mid + 1
		}
	}

	return -1
}
