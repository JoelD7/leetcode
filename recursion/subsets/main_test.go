package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsets(t *testing.T) {
	t.Run("nums = [1, 2, 3]", func(t *testing.T) {
		nums := []int{1, 2, 3}
		expected := [][]int{
			{},
			{1},
			{2},
			{1, 2},
			{3},
			{1, 3},
			{2, 3},
			{1, 2, 3},
		}
		assert.ElementsMatch(t, expected, subsets(nums))
	})

	t.Run("nums = [0] (Single element)", func(t *testing.T) {
		nums := []int{0}
		expected := [][]int{
			{},
			{0},
		}
		assert.ElementsMatch(t, expected, subsets(nums))
	})

	t.Run("nums = [1, 2] (Two elements)", func(t *testing.T) {
		nums := []int{1, 2}
		expected := [][]int{
			{},
			{1},
			{2},
			{1, 2},
		}
		assert.ElementsMatch(t, expected, subsets(nums))
	})

	t.Run("nums = [-1, -2] (Negative numbers)", func(t *testing.T) {
		nums := []int{-1, -2}
		expected := [][]int{
			{},
			{-1},
			{-2},
			{-1, -2},
		}
		assert.ElementsMatch(t, expected, subsets(nums))
	})

	t.Run("nums = [7, 8, 9, 10] (Four elements)", func(t *testing.T) {
		nums := []int{7, 8, 9, 10}
		expected := [][]int{
			{},
			{7}, {8}, {9}, {10},
			{7, 8}, {7, 9}, {7, 10}, {8, 9}, {8, 10}, {9, 10},
			{7, 8, 9}, {7, 8, 10}, {7, 9, 10}, {8, 9, 10},
			{7, 8, 9, 10},
		}
		assert.ElementsMatch(t, expected, subsets(nums))
	})
}
