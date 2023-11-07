package contains_duplicate

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	c := require.New(t)

	c.True(containsDuplicate([]int{1, 2, 3, 1}))
	c.True(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))

	c.False(containsDuplicate([]int{1, 2, 3, 4}))
}
