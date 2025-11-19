package permutations

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	visited := make(map[int]struct{})
	comb := make([]int, 0, len(nums))

	res = backtrack(nums, comb, res, visited)

	return res
}

func backtrack(nums, comb []int, res [][]int, visited map[int]struct{}) [][]int {
	if len(nums) == len(comb) {
		c := make([]int, len(comb))
		copy(c, comb)
		res = append(res, c)
		return res
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := visited[i]; ok {
			continue
		}

		comb = append(comb, nums[i])
		visited[i] = struct{}{}

		res = backtrack(nums, comb, res, visited)

		comb = comb[:len(comb)-1]
		delete(visited, i)
	}

	return res
}
