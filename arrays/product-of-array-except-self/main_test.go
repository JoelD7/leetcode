package product_of_array_except_self

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	c := require.New(t)

	result := productExceptSelf([]int{1, 2, 3, 4})

	c.Len(result, 4)
	c.Equal(24, result[0])
	c.Equal(12, result[1])
	c.Equal(8, result[2])
	c.Equal(6, result[3])
}

func TestPrefixProduct(t *testing.T) {
	c := require.New(t)

	result := prefixProduct([]int{1, 2, 3, 4})
	c.Len(result, 4)
	c.Equal(1, result[0])
	c.Equal(2, result[1])
	c.Equal(6, result[2])
	c.Equal(24, result[3])
}

func TestSuffixProduct(t *testing.T) {
	c := require.New(t)

	result := suffixProduct([]int{1, 2, 3, 4})
	c.Len(result, 4)
	c.Equal(24, result[0])
	c.Equal(24, result[1])
	c.Equal(12, result[2])
	c.Equal(4, result[3])
}
