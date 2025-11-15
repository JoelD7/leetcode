package word_search

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExists(t *testing.T) {
	c := require.New(t)

	t.Run(`board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"`, func(t *testing.T) {
		board := [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}

		c.True(exist(board, "ABCCED"))
	})

	t.Run(`board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"`, func(t *testing.T) {
		board := [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}

		c.True(exist(board, "SEE"))
	})

	t.Run(`board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"`, func(t *testing.T) {
		board := [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}

		c.False(exist(board, "ABCB"))
	})

	t.Run(`board = [["a"]], word = "a"`, func(t *testing.T) {
		board := [][]byte{
			{'a'},
		}

		c.True(exist(board, "a"))
	})

	t.Run(`board = [["a", "b"]], word = "ba"`, func(t *testing.T) {
		board := [][]byte{
			{'a', 'b'},
		}

		c.True(exist(board, "ba"))
	})
}
