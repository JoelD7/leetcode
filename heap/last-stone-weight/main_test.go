package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastStoneWeight(t *testing.T) {
	t.Run("multiple stones with a remaining weight", func(t *testing.T) {
		assert.Equal(t, 1, lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
	})

	t.Run("single stone input", func(t *testing.T) {
		assert.Equal(t, 1, lastStoneWeight([]int{1}))
	})

	t.Run("all stones destroyed completely", func(t *testing.T) {
		assert.Equal(t, 0, lastStoneWeight([]int{2, 2}))
	})

	t.Run("two stones with different weights", func(t *testing.T) {
		assert.Equal(t, 2, lastStoneWeight([]int{3, 5}))
	})
}
