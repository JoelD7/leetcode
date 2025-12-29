package evaluate_reverse_polish_notation

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	c := require.New(t)

	t.Run("2 1 + 3 *", func(t *testing.T) {
		c.Equal(9, evalRPN([]string{"2", "1", "+", "3", "*"}))
	})

	t.Run("4 13 5 / +", func(t *testing.T) {
		c.Equal(6, evalRPN([]string{"4", "13", "5", "/", "+"}))
	})

	t.Run("10 6 9 3 + -11 * / * 17 + 5 +", func(t *testing.T) {
		c.Equal(22, evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
	})
}
