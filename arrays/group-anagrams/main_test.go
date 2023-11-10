package group_anagrams

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	c := require.New(t)

	result := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	c.Len(result, 3)
}
