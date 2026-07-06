package main

func pacificAtlantic(heights [][]int) [][]int {
	dirs := [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	result := make([][]int, 0)
	pacific := make([][]bool, len(heights))
	atlantic := make([][]bool, len(heights))

	for i := 0; i < len(heights); i++ {
		pacific[i] = make([]bool, len(heights[0]))
		atlantic[i] = make([]bool, len(heights[0]))
	}

	var dfs func(r, c int, visited [][]bool)
	dfs = func(r, c int, visited [][]bool) {
		if visited[r][c] {
			return
		}

		visited[r][c] = true

		for _, dir := range dirs {
			x := r + dir[0]
			y := c + dir[1]

			if x < 0 || y < 0 || x >= len(heights) || y >= len(heights[r]) {
				continue
			}

			if heights[x][y] < heights[r][c] {
				continue
			}

			dfs(x, y, visited)
		}
	}

	for i := 0; i < len(heights); i++ {
		dfs(i, 0, pacific)
		dfs(i, len(heights[0])-1, atlantic)
	}

	for j := 0; j < len(heights[0]); j++ {
		dfs(0, j, pacific)
		dfs(len(heights)-1, j, atlantic)
	}

	for r := 0; r < len(heights); r++ {
		for c := 0; c < len(heights[r]); c++ {
			if pacific[r][c] && atlantic[r][c] {
				result = append(result, []int{r, c})
			}
		}
	}

	return result
}
