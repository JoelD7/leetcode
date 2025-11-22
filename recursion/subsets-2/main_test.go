package subsets_2

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSubsets(t *testing.T) {
	c := require.New(t)

	t.Run("nums = [1,2,2]", func(t *testing.T) {
		c.ElementsMatch([][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}, subsetsWithDup([]int{1, 2, 2}))
	})

	t.Run("nums = [0]", func(t *testing.T) {
		c.ElementsMatch([][]int{{}, {0}}, subsetsWithDup([]int{0}))
	})
}
