package search_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	l, r := 0, (m*n)-1

	for l <= r {
		mid := (l + r) / 2
		row := mid / n
		col := mid % n

		midVal := matrix[row][col]

		if target == midVal {
			return true
		} else if target > midVal {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}
