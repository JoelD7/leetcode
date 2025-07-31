package repeated_dna_sequences

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindRepeatedDNASequences(t *testing.T) {
	c := require.New(t)

	t.Run("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", func(t *testing.T) {
		c.Equal([]string{"AAAAACCCCC", "CCCCCAAAAA"}, findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	})

	t.Run("AAAAAAAAAAAAA", func(t *testing.T) {
		c.Equal([]string{"AAAAAAAAAA"}, findRepeatedDnaSequences("AAAAAAAAAAAAA"))
	})

	t.Run("AAAAAAAAAAA", func(t *testing.T) {
		c.Equal([]string{"AAAAAAAAAA"}, findRepeatedDnaSequences("AAAAAAAAAAA"))
	})
}
