package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	first, second := cost[0], cost[1]
	var cur int
	for i := 2; i < n; i++ {
		cur = cost[i] + min(first, second)
		first = second
		second = cur
	}

	return min(first, second)
}
