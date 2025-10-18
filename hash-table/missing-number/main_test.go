package missing_number

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	c := require.New(t)

	t.Run("[3,0,1]", func(t *testing.T) {
		c.Equal(2, missingNumber([]int{3, 0, 1}))
	})

	t.Run("[0,1]", func(t *testing.T) {
		c.Equal(2, missingNumber([]int{0, 1}))
	})

	t.Run("[9,6,4,2,3,5,7,0,1]", func(t *testing.T) {
		c.Equal(8, missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	})

	t.Run("[1]", func(t *testing.T) {
		c.Equal(0, missingNumber([]int{1}))
	})
}
