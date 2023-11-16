package valid_sudoku

// 1. Each row must contain the digits 1-9 without repetition.
// 2. Each column must contain the digits 1-9 without repetition.
// 3. Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
func isValidSudoku(board [][]byte) bool {
	rowMap := map[byte]struct{}{}
	colMap := map[byte]struct{}{}
	var rowVal byte
	var colVal byte

	//key(r+c)=cuadrant, value=digits in that cuadrant
	tbtMap := map[string]map[byte]struct{}{
		"00": {},
		"01": {},
		"02": {},
		"10": {},
		"11": {},
		"12": {},
		"20": {},
		"21": {},
		"22": {},
	}

	for row := 0; row < 9; row++ {
		rowMap = map[byte]struct{}{}
		colMap = map[byte]struct{}{}

		for col := 0; col < 9; col++ {
			rowVal = board[row][col]
			colVal = board[col][row]

			c := rune(rowVal)
			if c >= '0' && c <= '9' {
				_, ok := rowMap[rowVal]
				if ok {
					return false
				}

				rowMap[rowVal] = struct{}{}
				cuadrant := getCuadrant(row, col)
				_, ok = tbtMap[cuadrant][rowVal]
				if ok {
					return false
				}

				tbtMap[cuadrant][rowVal] = struct{}{}

			}

			c = rune(colVal)
			if c >= '0' && c <= '9' {
				_, ok := colMap[colVal]
				if ok {
					return false
				}

				colMap[colVal] = struct{}{}
			}
		}
	}

	return true
}

func getCuadrant(r, c int) string {
	var row, col string
	rowDiff := r - 3
	colDiff := c - 3

	if rowDiff < 0 {
		row = "0"
	}

	if rowDiff >= 0 && rowDiff <= 2 {
		row = "1"
	}

	if rowDiff >= 3 && rowDiff <= 5 {
		row = "2"
	}

	if colDiff < 0 {
		col = "0"
	}

	if colDiff >= 0 && colDiff <= 2 {
		col = "1"
	}

	if colDiff >= 3 && colDiff <= 5 {
		col = "2"
	}

	return row + col
}
