package word_search

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			//No way to form word
			if board[i][j] != word[0] {
				continue
			}

			if backtrack(i, j, 0, board, visited, word) {
				return true
			}
		}
	}

	return false
}

func backtrack(i, j, index int, board [][]byte, visited [][]bool, word string) bool {
	if index == len(word) {
		return true
	}

	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || visited[i][j] || board[i][j] != word[index] {
		return false
	}

	visited[i][j] = true

	if backtrack(i-1, j, index+1, board, visited, word) ||
		backtrack(i, j-1, index+1, board, visited, word) ||
		backtrack(i, j+1, index+1, board, visited, word) ||
		backtrack(i+1, j, index+1, board, visited, word) {
		return true
	}

	visited[i][j] = false

	return false
}
