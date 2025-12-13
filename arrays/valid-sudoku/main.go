package valid_sudoku

func isValidSudoku(board [][]byte) bool {
	//Maps row number to found digits
	rowDigits := map[int]map[byte]bool{
		0: make(map[byte]bool),
		1: make(map[byte]bool),
		2: make(map[byte]bool),
		3: make(map[byte]bool),
		4: make(map[byte]bool),
		5: make(map[byte]bool),
		6: make(map[byte]bool),
		7: make(map[byte]bool),
		8: make(map[byte]bool),
	}

	//Maps col number to found digits
	colDigits := map[int]map[byte]bool{
		0: make(map[byte]bool),
		1: make(map[byte]bool),
		2: make(map[byte]bool),
		3: make(map[byte]bool),
		4: make(map[byte]bool),
		5: make(map[byte]bool),
		6: make(map[byte]bool),
		7: make(map[byte]bool),
		8: make(map[byte]bool),
	}

	subGridDigits := map[int]map[int]map[byte]bool{
		0: {
			0: make(map[byte]bool),
			1: make(map[byte]bool),
			2: make(map[byte]bool),
		},
		1: {
			0: make(map[byte]bool),
			1: make(map[byte]bool),
			2: make(map[byte]bool),
		},
		2: {
			0: make(map[byte]bool),
			1: make(map[byte]bool),
			2: make(map[byte]bool),
		},
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				continue
			}

			if rowDigits[i][board[i][j]] {
				return false
			}

			rowDigits[i][board[i][j]] = true

			if colDigits[j][board[i][j]] {
				return false
			}

			colDigits[j][board[i][j]] = true

			if subGridDigits[i/3][j/3][board[i][j]] {
				return false
			}

			subGridDigits[i/3][j/3][board[i][j]] = true
		}
	}

	return true
}
