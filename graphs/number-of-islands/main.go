package main

func numIslands(grid [][]byte) int {
	islands := 0

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0'
		dfs(r, c-1)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r+1, c)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != '0' {
				islands++
				grid[i][j] = '0'

				dfs(i, j-1)
				dfs(i-1, j)
				dfs(i, j+1)
				dfs(i+1, j)
			}
		}
	}

	return islands
}
