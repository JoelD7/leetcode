package search_rotated_sorted_array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearch(t *testing.T) {
	c := require.New(t)

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
}
