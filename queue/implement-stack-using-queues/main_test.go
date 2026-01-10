package implement_stack_using_queues

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMainF(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		stack := Constructor()
		stack.Push(1)
		stack.Push(2)
		assert.Equal(t, 2, stack.Top()) // return 2
		assert.Equal(t, 2, stack.Pop()) // return 2
		assert.False(t, stack.Empty())
	})

	t.Run("Test case 2", func(t *testing.T) {
		stack := Constructor()
		stack.Push(1)
		assert.Equal(t, 1, stack.Pop())
		assert.True(t, stack.Empty())
	})

	//["MyStack","push","push","pop","top"]
	//[[],[1],[2],[],[]]
	t.Run("Test case 3", func(t *testing.T) {
		stack := Constructor()
		stack.Push(1)
		stack.Push(2)
		assert.Equal(t, 2, stack.Pop())
		assert.Equal(t, 1, stack.Top())
	})

}
