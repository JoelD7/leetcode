package three_sum_closest

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	c := require.New(t)

	t.Run("-1,2,1,-4, target = 1", func(t *testing.T) {
		c.Equal(2, threeSumClosest([]int{-1, 2, 1, -4}, 1))
	})

	t.Run("[0,0,0], target = 1", func(t *testing.T) {
		c.Equal(0, threeSumClosest([]int{0, 0, 0}, 1))
	})
}
