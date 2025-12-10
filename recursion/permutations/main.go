package permutations

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	comb := make([]int, 0)
	visited := make(map[int]bool)

	var backtrack func()

	backtrack = func() {
		if len(comb) == len(nums) {
			res = append(res, append([]int(nil), comb...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}

			visited[i] = true
			comb = append(comb, nums[i])
			backtrack()

			comb = comb[:len(comb)-1]
			visited[i] = false
		}
	}

	backtrack()
	return res
}
