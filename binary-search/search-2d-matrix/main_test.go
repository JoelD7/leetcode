package search_2d_matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	t.Run("Large matrix", func(t *testing.T) {
		matrix := [][]int{
			{-8, -7, -5, -3, -3, -1, 1},
			{2, 2, 2, 3, 3, 5, 7},
			{8, 9, 11, 11, 13, 15, 17},
			{18, 18, 18, 20, 20, 20, 21},
			{23, 24, 26, 26, 26, 27, 27},
			{28, 29, 29, 30, 32, 32, 34},
		}
		assert.True(t, searchMatrix(matrix, -5))
	})
}
