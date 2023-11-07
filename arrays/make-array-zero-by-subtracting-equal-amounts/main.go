package make_array_zero_by_subtracting_equal_amounts

import "sort"

func minimumOperations(nums []int) int {
	x := 0
	operationsCount := 0

	for {
		sort.Ints(nums)

		for _, val := range nums {
			if val != 0 && val > 1 {
				x = val - 1
			}

			if val != 0 {
				x = val
			}
		}

		if x == 0 && operationsCount == 0 {
			return 0
		}

		if x == 0 {
			return operationsCount
		}

		for i, val := range nums {
			if val != 0 {
				nums[i] = val - x
			}
		}

		operationsCount++
		x = 0
	}
}
