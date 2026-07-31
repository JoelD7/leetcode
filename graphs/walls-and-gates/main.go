package main

func islandsAndTreasure(grid [][]int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return
	}

	m := len(grid)
	n := len(grid[0])
	queue := make([][2]int, 0)

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 0 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	dirs := [][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		r, c := node[0], node[1]

		for _, dir := range dirs {
			x, y := r+dir[0], c+dir[1]

			if x < 0 || y < 0 || x >= m || y >= n || grid[x][y] != 2147483647 {
				continue
			}

			grid[x][y] = grid[r][c] + 1
			queue = append(queue, [2]int{x, y})
		}
	}
}
