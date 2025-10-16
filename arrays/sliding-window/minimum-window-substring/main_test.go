package minimum_window_substring

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMinWindow(t *testing.T) {
	c := require.New(t)

	t.Run("s = ADOBECODEBANC, t = ABC", func(t *testing.T) {
		c.Equal("BANC", minWindow("ADOBECODEBANC", "ABC"))
	})

	t.Run("s = a, t = a", func(t *testing.T) {
		c.Equal("a", minWindow("a", "a"))
	})

	t.Run("s = a, t = aa", func(t *testing.T) {
		c.Equal("", minWindow("a", "aa"))
	})
}
