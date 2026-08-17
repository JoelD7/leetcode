package main

func numDecodings(s string) int {
	dp := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = -1
	}

	var decodeWays func(i int) int
	decodeWays = func(i int) int {
		if i == len(s) {
			return 1
		}

		if s[i] == '0' {
			return 0
		}

		if dp[i] != -1 {
			return dp[i]
		}

		one := decodeWays(i + 1)
		two := 0

		if i+1 < len(s) {
			num := (int(s[i]-'0') * 10) + (int(s[i+1] - '0'))
			if num <= 26 {
				two = decodeWays(i + 2)
			}
		}

		dp[i] = one + two
		return dp[i]
	}

	return decodeWays(0)
}
