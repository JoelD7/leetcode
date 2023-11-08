package valid_anagram

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	c := require.New(t)

	c.True(isAnagram("anagram", "nagaram"))
}
