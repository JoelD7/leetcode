package set_matrix_zeroes

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	c := require.New(t)

	t.Run("[[1,1,1],[1,0,1],[1,1,1]]", func(t *testing.T) {
		matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
		setZeroes(matrix)
		c.Equal([][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}, matrix)
	})

	//	[[0,1,2,0],[3,4,5,2],[1,3,1,5]]
	t.Run("[[0,1,2,0],[3,4,5,2],[1,3,1,5]]", func(t *testing.T) {
		matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
		setZeroes(matrix)
		c.Equal([][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}, matrix)
	})
}
