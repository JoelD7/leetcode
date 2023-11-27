package three_sum

func threeSum(nums []int) [][]int {
	//prefixSum := getPrefixSum(nums)
	triplets := make([][]int, 0)

	var l, r int

	//{-1, 0, 1, 2, -1, -4}
	for i := 0; i < len(nums); i++ {
		l = 0
		r = len(nums) - 1

		for l < r {
			if l == i {
				l++
				continue
			}

			if r == i {
				r--
				continue
			}

			if nums[i]+nums[l]+nums[r] == 0 {
				triplets = append(triplets, []int{nums[i], nums[l], nums[r]})
			}

			l++
			r--
		}
	}

	return triplets
}

func getPrefixSum(nums []int) []int {
	prev := nums[0]
	ps := make([]int, 0)
	ps = append(ps, prev)

	for i := 1; i < len(nums); i++ {
		ps = append(ps, prev+nums[i])
		prev += nums[i]
	}

	return ps
}

func rangeSum(l, r int, prefixSum []int) int {
	if l == 0 {
		return prefixSum[r]
	}

	return prefixSum[r] - prefixSum[l-1]
}
