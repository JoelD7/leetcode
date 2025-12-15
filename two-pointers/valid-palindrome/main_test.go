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
