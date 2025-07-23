package contains_duplicate_2

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	c := require.New(t)

	t.Run("99,99 k = 2", func(t *testing.T) {
		c.True(containsNearbyDuplicate([]int{99, 99}, 2))
	})

	t.Run("[1,2,3,1], k=3", func(t *testing.T) {
		c.True(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	})
}
