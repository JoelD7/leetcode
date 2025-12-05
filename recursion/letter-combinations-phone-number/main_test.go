package letter_combinations_phone_number

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	c := require.New(t)

	t.Run("23", func(t *testing.T) {
		c.ElementsMatch([]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, letterCombinations("23"))
	})

	t.Run("2", func(t *testing.T) {
		c.ElementsMatch([]string{"a", "b", "c"}, letterCombinations("2"))
	})
}
