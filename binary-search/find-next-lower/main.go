package find_next_lower

//1,2,3,4,5,6,7,8,9,10

// 3,4
// for number = 4
func findNextLower(number int, arr []int) (int, int) {
	iterations := 0

	left, mid, right := 0, 0, len(arr)-1

	for left <= right {
		iterations++

		mid = left + (right-left)/2

		if arr[mid] >= number {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return arr[left-1], iterations
}
