package word_search

func exist(board [][]byte, word string) bool {
	var backtrack func(index, i, j int) bool
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[i]))
	}

	backtrack = func(index, i, j int) bool {
		if index == len(word) {
			return true
		}

		if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) || visited[i][j] {
			return false
		}

		if word[index] != board[i][j] {
			return false
		}

		visited[i][j] = true

		if backtrack(index+1, i-1, j) || backtrack(index+1, i+1, j) ||
			backtrack(index+1, i, j-1) || backtrack(index+1, i, j+1) {
			return true
		}

		visited[i][j] = false

		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != word[0] {
				continue
			}
			if backtrack(0, i, j) {
				return true
			}
		}
	}

	return false
}
