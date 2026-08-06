package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	t.Run("n = 2", func(t *testing.T) {
		assert.Equal(t, 2, climbStairs(2))
	})

	t.Run("n = 3", func(t *testing.T) {
		assert.Equal(t, 3, climbStairs(3))
	})

	t.Run("n = 44", func(t *testing.T) {
		assert.Equal(t, 1134903170, climbStairs(44))
	})
}
