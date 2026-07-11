package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFinish(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		assert.True(t, canFinish(2, [][]int{{1, 0}}))
	})

	t.Run("Example 2", func(t *testing.T) {
		assert.False(t, canFinish(2, [][]int{{1, 0}, {0, 1}}))
	})
}
