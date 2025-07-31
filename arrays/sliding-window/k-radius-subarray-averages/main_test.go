package k_radius_subarray_averages

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAverages(t *testing.T) {
	c := require.New(t)

	t.Run("[7,4,3,9,1,8,5,2,6], k = 3", func(t *testing.T) {
		c.Equal([]int{-1, -1, -1, 5, 4, 4, -1, -1, -1}, getAverages([]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3))
	})

	t.Run("[100000], k = 0", func(t *testing.T) {
		c.Equal([]int{100000}, getAverages([]int{100000}, 0))
	})

	t.Run("[8], k = 100000", func(t *testing.T) {
		c.Equal([]int{-1}, getAverages([]int{8}, 100000))
	})
}
