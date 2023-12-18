package evaluate_reverse_polish_notation

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	c := require.New(t)

	c.Equal(9, evalRPN([]string{"2", "1", "+", "3", "*"}))
	c.Equal(6, evalRPN([]string{"4", "13", "5", "/", "+"}))
	c.Equal(22, evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
