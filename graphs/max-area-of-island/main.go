package main

import (
	"fmt"
)

func maxAreaOfIsland(grid [][]int) int {
	var dfs func(r, c int)
	maxArea := 0
	count := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[r]) || grid[r][c] == 0 {
			return
		}

		count++
		grid[r][c] = 0

		dfs(r, c-1)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r+1, c)
	}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			count = 0
			if r == 3 && c == 8 {
				fmt.Println()
			}
			if grid[r][c] != 0 {
				count++
				grid[r][c] = 0

				dfs(r, c-1)
				dfs(r-1, c)
				dfs(r, c+1)
				dfs(r+1, c)

				maxArea = max(maxArea, count)
			}
		}
	}

	return maxArea
}
