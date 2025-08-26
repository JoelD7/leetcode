package longest_repeating_char_replacement

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCharacterReplacement(t *testing.T) {
	c := require.New(t)

	t.Run("ABAB, k = 2", func(t *testing.T) {
		c.Equal(4, characterReplacement("ABAB", 2))
	})

	t.Run("AABABBA, k = 1", func(t *testing.T) {
		c.Equal(4, characterReplacement("AABABBA", 1))
	})

	t.Run("BBBBBB, k = 1", func(t *testing.T) {
		c.Equal(6, characterReplacement("BBBBBB", 1))
	})

	t.Run("ABAA, k = 0", func(t *testing.T) {
		c.Equal(2, characterReplacement("ABAA", 0))
	})

	t.Run("AAAB, k = 0", func(t *testing.T) {
		c.Equal(3, characterReplacement("AAAB", 0))
	})

	t.Run("ABBB, k = 2", func(t *testing.T) {
		c.Equal(4, characterReplacement("ABBB", 2))
	})

	t.Run("BAAAB, k = 2", func(t *testing.T) {
		c.Equal(5, characterReplacement("BAAAB", 2))
	})
}
