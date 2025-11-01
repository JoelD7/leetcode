package combinations

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCombine(t *testing.T) {
	c := require.New(t)

	t.Run("n = 4, k = 2", func(t *testing.T) {
		c.Equal([][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}, combine(4, 2))
	})

	t.Run("n = 1, k = 1", func(t *testing.T) {
		c.Equal([][]int{{1}}, combine(1, 1))
	})
}
