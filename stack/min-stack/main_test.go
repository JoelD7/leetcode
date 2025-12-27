package min_stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMinStack(t *testing.T) {
	c := require.New(t)

	t.Run("Example 1", func(t *testing.T) {
		minStack := Constructor()
		minStack.Push(-2)
		minStack.Push(0)
		minStack.Push(-3)
		c.Equal(-3, minStack.GetMin())
		minStack.Pop()
		c.Equal(0, minStack.Top())
		c.Equal(-2, minStack.GetMin())
	})

	t.Run("Test case 2", func(t *testing.T) {
		minStack := Constructor()
		//   ["MinStack","push","push","push","top","pop","getMin","pop","getMin","pop","push","top", "getMin","push","top","getMin","pop","getMin"]
		//	 [[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]
		minStack.Push(2147483646)
		minStack.Push(2147483646)
		minStack.Push(2147483647)
		c.Equal(2147483647, minStack.Top())
		minStack.Pop()
		c.Equal(2147483646, minStack.GetMin())
		minStack.Pop()
		c.Equal(2147483646, minStack.GetMin())
		minStack.Pop()
		minStack.Push(2147483647)
		c.Equal(2147483647, minStack.Top())
		c.Equal(2147483647, minStack.GetMin())
		minStack.Push(-2147483648)
		c.Equal(-2147483648, minStack.Top())
		c.Equal(-2147483648, minStack.GetMin())
		minStack.Pop()
		c.Equal(2147483647, minStack.GetMin())
	})

	t.Run("Test case 3", func(t *testing.T) {
		minStack := Constructor()
		minStack.Push(0)
		minStack.Push(1)
		minStack.Push(0)
		c.Equal(0, minStack.GetMin())
		minStack.Pop()
		c.Equal(0, minStack.GetMin())
	})
}
