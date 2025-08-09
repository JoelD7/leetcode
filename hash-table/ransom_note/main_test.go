package ransom_note

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	c := require.New(t)

	t.Run("ransomNote = a, magazine = b", func(t *testing.T) {
		c.False(canConstruct("a", "b"))
	})

	t.Run("ransomNote = aa, magazine = ab", func(t *testing.T) {
		c.False(canConstruct("aa", "ab"))
	})

	t.Run("ransomNote = aa, magazine = aab", func(t *testing.T) {
		c.True(canConstruct("aa", "aab"))
	})
}
