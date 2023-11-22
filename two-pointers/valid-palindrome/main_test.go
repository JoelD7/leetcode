package valid_palindrome

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	c := require.New(t)

	c.True(isPalindrome("aa"))
	c.False(isPalindrome("op"))
	c.True(isPalindrome("A man, a plan, a canal: Panama"))
}

func TestNormalize(t *testing.T) {
	c := require.New(t)

	c.Equal("amanaplanacanalpanama", normalize("A man, a plan, a canal: Panama"))
	c.Equal("raceacar", normalize("race a car"))
	c.Equal("abrigo", normalize("?4ab8r)-i@#g== 0o"))
}
