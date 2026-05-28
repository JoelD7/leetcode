package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeastInterval(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		assert.Equal(t, 8, leastInterval([]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2))
	})
}
