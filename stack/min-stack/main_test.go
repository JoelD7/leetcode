package min_stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMinStack(t *testing.T) {
	c := require.New(t)

	build := Constructor()

	minStack := &build
	minStack.Push(0)
	minStack.Push(1)
	minStack.Push(0)
	c.Equal(0, minStack.GetMin())
	minStack.Pop()
	c.Equal(0, minStack.GetMin())
	//
	build = Constructor()
	minStack = &build

	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	c.Equal(-3, minStack.GetMin())
	minStack.Pop()
	c.Equal(0, minStack.Top())
	c.Equal(-2, minStack.GetMin())

}
