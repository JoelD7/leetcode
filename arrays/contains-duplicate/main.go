package contains_duplicate

func containsDuplicate(nums []int) bool {
	seen := map[int]struct{}{}

	for _, val := range nums {
		_, ok := seen[val]
		if ok {
			return true
		}

		seen[val] = struct{}{}
	}

	return false
}
