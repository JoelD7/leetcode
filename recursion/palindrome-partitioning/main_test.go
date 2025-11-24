package palindrome_partitioning

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPartition(t *testing.T) {
	c := require.New(t)

	t.Run("aab", func(t *testing.T) {
		c.Equal([][]string{{"a", "a", "b"}, {"aa", "b"}}, partition("aab"))
	})

	t.Run("a", func(t *testing.T) {
		c.Equal([][]string{{"a"}}, partition("a"))
	})
}
