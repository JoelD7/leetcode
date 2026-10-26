package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetZeroes(t *testing.T) {
	t.Run("single zero in the center", func(t *testing.T) {
		matrix := [][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		}
		expected := [][]int{
			{1, 0, 1},
			{0, 0, 0},
			{1, 0, 1},
		}
		setZeroes(matrix)
		assert.Equal(t, expected, matrix)
	})

	t.Run("zeros on the edges", func(t *testing.T) {
		matrix := [][]int{
			{0, 1, 2, 0},
			{3, 4, 5, 2},
			{1, 3, 1, 5},
		}
		expected := [][]int{
			{0, 0, 0, 0},
			{0, 4, 5, 0},
			{0, 3, 1, 0},
		}
		setZeroes(matrix)
		assert.Equal(t, expected, matrix)
	})

	t.Run("matrix with no zeros", func(t *testing.T) {
		matrix := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
		expected := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
		setZeroes(matrix)
		assert.Equal(t, expected, matrix)
	})

	t.Run("matrix with all zeros", func(t *testing.T) {
		matrix := [][]int{
			{0, 0},
			{0, 0},
		}
		expected := [][]int{
			{0, 0},
			{0, 0},
		}
		setZeroes(matrix)
		assert.Equal(t, expected, matrix)
	})

	t.Run("single row matrix", func(t *testing.T) {
		matrix := [][]int{
			{1, 0, 3},
		}
		expected := [][]int{
			{0, 0, 0},
		}
		setZeroes(matrix)
		assert.Equal(t, expected, matrix)
	})

	t.Run("single column matrix", func(t *testing.T) {
		matrix := [][]int{
			{1},
			{0},
			{3},
		}
		expected := [][]int{
			{0},
			{0},
			{0},
		}
		setZeroes(matrix)
		assert.Equal(t, expected, matrix)
	})
}
