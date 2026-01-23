package search_rotated_sorted_array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearch(t *testing.T) {
	c := require.New(t)

	t.Run("4,5,6,7,0,1,2", func(t *testing.T) {
		c.Equal(4, search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	})

	t.Run("[1,3], target: 3", func(t *testing.T) {
		c.Equal(1, search([]int{1, 3}, 3))
	})

	t.Run("[1,3,5], target: 5", func(t *testing.T) {
		c.Equal(2, search([]int{1, 3, 5}, 5))
	})

	t.Run("[5,1,3], target: 3", func(t *testing.T) {
		c.Equal(2, search([]int{5, 1, 3}, 3))
	})

	t.Run("[5,1,2,3,4], target: 1", func(t *testing.T) {
		c.Equal(1, search([]int{5, 1, 2, 3, 4}, 1))
	})

	t.Run("[1,3], target: 0", func(t *testing.T) {
		c.Equal(-1, search([]int{1, 3}, 0))
	})

	t.Run("[1], target: 0", func(t *testing.T) {
		c.Equal(-1, search([]int{1}, 0))
	})

	t.Run("[3,1], target: 1", func(t *testing.T) {
		c.Equal(1, search([]int{3, 1}, 1))
	})

	t.Run("[3,5,1], target: 3", func(t *testing.T) {
		c.Equal(0, search([]int{3, 5, 1}, 3))
	})

	t.Run("[5,1,3], target: 5", func(t *testing.T) {
		c.Equal(0, search([]int{5, 1, 3}, 5))
	})

}
