package search_2d_matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	t.Run("Target is present in the first row", func(t *testing.T) {
		matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
		target := 3
		assert.Equal(t, true, searchMatrix(matrix, target))
	})

	t.Run("Target is missing from the matrix", func(t *testing.T) {
		matrix := [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		}
		target := 13
		assert.Equal(t, false, searchMatrix(matrix, target))
	})

	t.Run("Target is the first element of a middle row", func(t *testing.T) {
		matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
		target := 10
		assert.Equal(t, true, searchMatrix(matrix, target))
	})

	t.Run("Target is the last element of the last row", func(t *testing.T) {
		matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
		target := 60
		assert.Equal(t, true, searchMatrix(matrix, target))
	})

	t.Run("Matrix has only one element and it is the target", func(t *testing.T) {
		matrix := [][]int{{1}}
		target := 1
		assert.Equal(t, true, searchMatrix(matrix, target))
	})

	t.Run("Matrix has only one element and it is not the target", func(t *testing.T) {
		matrix := [][]int{{1}}
		target := 0
		assert.Equal(t, false, searchMatrix(matrix, target))
	})

	t.Run("Target is smaller than the smallest element in the matrix", func(t *testing.T) {
		matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
		target := 0
		assert.Equal(t, false, searchMatrix(matrix, target))
	})

	t.Run("Target is larger than the largest element in the matrix", func(t *testing.T) {
		matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
		target := 100
		assert.Equal(t, false, searchMatrix(matrix, target))
	})

	t.Run("Matrix is a single column", func(t *testing.T) {
		matrix := [][]int{{1}, {3}, {5}}
		target := 3
		assert.Equal(t, true, searchMatrix(matrix, target))
	})
}
