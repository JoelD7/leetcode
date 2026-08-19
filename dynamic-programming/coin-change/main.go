package main

func coinChange(coins []int, amount int) int {
	MaxInt := amount + 1
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = MaxInt
	}

	var min func(a, b int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp[0] = 0
	var rem int

	for curAmount := 1; curAmount <= amount; curAmount++ {
		for _, coin := range coins {
			rem = curAmount - coin

			if rem >= 0 {
				dp[curAmount] = min(dp[curAmount], dp[rem]+1)
			}
		}
	}

	if dp[amount] == MaxInt {
		return -1
	}

	return dp[amount]
}
