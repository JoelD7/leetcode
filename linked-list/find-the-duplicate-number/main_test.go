package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicate(t *testing.T) {
	t.Run("duplicate is 2 in unsorted array", func(t *testing.T) {
		assert.Equal(t, 2, findDuplicate([]int{1, 3, 4, 2, 2}))
	})

	t.Run("duplicate is 3 scattered", func(t *testing.T) {
		assert.Equal(t, 3, findDuplicate([]int{3, 1, 3, 4, 2}))
	})

	t.Run("all elements are the same duplicate", func(t *testing.T) {
		assert.Equal(t, 3, findDuplicate([]int{3, 3, 3, 3, 3}))
	})

	t.Run("duplicate is 1 in small array", func(t *testing.T) {
		assert.Equal(t, 1, findDuplicate([]int{1, 1, 2}))
	})

	t.Run("duplicate at the very end", func(t *testing.T) {
		assert.Equal(t, 9, findDuplicate([]int{2, 5, 9, 6, 9, 3, 8, 7, 1, 9}))
		assert.Equal(t, 5, findDuplicate([]int{2, 1, 3, 4, 5, 5}))
	})
}
