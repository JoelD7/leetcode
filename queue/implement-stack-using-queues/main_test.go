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

	t.Run("Test case 3", func(t *testing.T) {
		stack := Constructor()
		stack.Push(1)
		stack.Push(2)
		assert.Equal(t, 2, stack.Pop())
		assert.Equal(t, 1, stack.Top())
	})

}

func TestQueue(t *testing.T) {
	q := NewQueue()

	assert.True(t, q.IsEmpty())
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	assert.Equal(t, 4, q.size)
	assert.False(t, q.IsEmpty())
	assert.Equal(t, 1, q.Peek())
	assert.Equal(t, 4, q.size)

	assert.Equal(t, 1, q.Dequeue())
	assert.Equal(t, 3, q.size)

	assert.Equal(t, 2, q.Peek())
	assert.Equal(t, 3, q.size)

	q.Enqueue(5)
	q.Enqueue(6)
	q.Enqueue(7)
	assert.Equal(t, 6, q.size)
	assert.Equal(t, 2, q.Dequeue())
	assert.Equal(t, 5, q.size)
	assert.Equal(t, 3, q.Dequeue())
	assert.Equal(t, 4, q.size)

	assert.Equal(t, 4, q.Dequeue())
	assert.Equal(t, 3, q.size)

	assert.Equal(t, 5, q.Peek())
	assert.Equal(t, 5, q.Dequeue())
	assert.Equal(t, 6, q.Dequeue())
	assert.Equal(t, 7, q.Dequeue())

	assert.Equal(t, -1, q.Dequeue())
}
