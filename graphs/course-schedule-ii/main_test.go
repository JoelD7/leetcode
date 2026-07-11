package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindOrder(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		prerequisites := [][]int{{1, 0}}
		assert.ElementsMatch(t, []int{0, 1}, findOrder(2, prerequisites))
	})

	t.Run("Example 2", func(t *testing.T) {
		prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
		assert.ElementsMatch(t, []int{0, 1, 2, 3}, findOrder(4, prerequisites))
	})

	t.Run("Example 3", func(t *testing.T) {
		prerequisites := [][]int{}
		assert.ElementsMatch(t, []int{0}, findOrder(1, prerequisites))
	})
}
