package two_sum

func twoSum(nums []int, target int) []int {
	indexByDif := make(map[int]int)

	for i, val := range nums {
		index, ok := indexByDif[val]
		// Check ok here to prevent false negatives, meaning, if a value doesn't exist "index" will be zero
		if i == index && ok {
			continue
		}

		if ok {
			return []int{i, index}
		}

		dif := target - val

		indexByDif[dif] = i
	}

	return []int{0, 0}
}
