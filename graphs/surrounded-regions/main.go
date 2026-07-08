package main

func solve(board [][]byte) {
	resilient := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		resilient[i] = make([]int, len(board[0]))
	}

	var dfs func(r, c int)

	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= len(board) || c >= len(board[r]) || board[r][c] == 'X' || resilient[r][c] == 1 {
			return
		}

		resilient[r][c] = 1
		dfs(r, c-1)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r+1, c)
	}

	//Left and right edges
	for r := 0; r < len(board); r++ {
		if board[r][0] == 'O' {
			dfs(r, 0)
		}

		if board[r][len(board[0])-1] == 'O' {
			dfs(r, len(board[0])-1)
		}
	}

	//Top and bottom edges
	for c := 0; c < len(board[0]); c++ {
		if board[0][c] == 'O' {
			dfs(0, c)
		}

		if board[len(board)-1][c] == 'O' {
			dfs(len(board)-1, c)
		}
	}

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[r]); c++ {
			if resilient[r][c] == 0 {
				board[r][c] = 'X'
			}
		}
	}
}
