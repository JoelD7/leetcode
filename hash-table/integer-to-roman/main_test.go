package integer_to_roman

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	c := require.New(t)

	t.Run("3749", func(t *testing.T) {
		c.Equal("MMMDCCXLIX", intToRoman(3749))
	})

	t.Run("58", func(t *testing.T) {
		c.Equal("LVIII", intToRoman(58))
	})

	t.Run("1994", func(t *testing.T) {
		c.Equal("MCMXCIV", intToRoman(1994))
	})

	t.Run("101", func(t *testing.T) {
		c.Equal("CI", intToRoman(101))
	})
}

func BenchmarkIntToRoman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intToRoman(3749)
		intToRoman(58)
		intToRoman(1994)
		intToRoman(101)
	}
}

func BenchmarkIntToRomanRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intToRomanRecursive(3749)
		intToRomanRecursive(58)
		intToRomanRecursive(1994)
		intToRomanRecursive(101)
	}
}
