package search_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	//0 = row, 1 = column
	min := []int{0, 0}
	max := []int{m - 1, n - 1}

	var r, c, middle, minRow, maxRow int
	for min[0] <= max[0] && min[1] <= max[1] {
		r = min[0] + (max[0]-min[0])/2
		c = min[1] + (max[1]-min[1])/2
		middle = matrix[r][c]

		minRow = matrix[r][0]
		maxRow = matrix[r][n-1]

		if middle == target || target == minRow || target == maxRow {
			return true
		}

		if target < minRow {
			max = []int{r - 1, n - 1}
			continue
		}

		if target > maxRow {
			min = []int{r + 1, 0}
			continue
		}

		if middle < target {
			min = []int{r, c + 1}
			max = []int{r, n - 1}
			continue
		}

		if middle > target {
			min[0] = r
			max = []int{r, c - 1}
			continue
		}
	}

	return false
}
