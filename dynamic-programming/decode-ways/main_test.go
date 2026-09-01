package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumDecodings(t *testing.T) {
	t.Run("basic example 1: 12", func(t *testing.T) {
		assert.Equal(t, 2, numDecodings("12"))
	})

	t.Run("basic example 2: 226", func(t *testing.T) {
		assert.Equal(t, 3, numDecodings("226"))
	})

	t.Run("leading zero: 06", func(t *testing.T) {
		assert.Equal(t, 0, numDecodings("06"))
	})

	t.Run("ending with 10", func(t *testing.T) {
		assert.Equal(t, 1, numDecodings("10"))
	})

	t.Run("invalid sequence: 30", func(t *testing.T) {
		assert.Equal(t, 0, numDecodings("30"))
	})

	t.Run("single digit: 1", func(t *testing.T) {
		assert.Equal(t, 1, numDecodings("1"))
	})
}
