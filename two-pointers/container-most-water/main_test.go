package container_most_water

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaxArea(t *testing.T) {
	c := require.New(t)

	c.Equal(17, maxArea([]int{2, 3, 4, 5, 18, 17, 6}))
	c.Equal(49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

}
