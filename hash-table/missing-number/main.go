package missing_number

func missingNumber(nums []int) int {
	n := len(nums)
	totalSum := (n * (n + 1)) / 2

	numsSum := 0
	for _, num := range nums {
		numsSum += num
	}

	return totalSum - numsSum
}
