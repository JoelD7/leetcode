package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFinish(t *testing.T) {
	t.Run("single course with no prerequisites", func(t *testing.T) {
		assert.Equal(t, true, canFinish(1, [][]int{}))
	})

	t.Run("two courses with valid linear prerequisite", func(t *testing.T) {
		assert.Equal(t, true, canFinish(2, [][]int{{1, 0}}))
	})

	t.Run("two courses with a direct cycle", func(t *testing.T) {
		assert.Equal(t, false, canFinish(2, [][]int{{1, 0}, {0, 1}}))
	})

	t.Run("multiple courses with a valid topological order", func(t *testing.T) {
		assert.Equal(t, true, canFinish(4, [][]int{{1, 0}, {2, 1}, {3, 2}}))
	})

	t.Run("multiple courses with a hidden cycle", func(t *testing.T) {
		assert.Equal(t, false, canFinish(4, [][]int{{1, 0}, {2, 1}, {3, 2}, {1, 3}}))
	})

	t.Run("disconnected graph with no cycles", func(t *testing.T) {
		assert.Equal(t, true, canFinish(3, [][]int{{1, 0}}))
	})
}
