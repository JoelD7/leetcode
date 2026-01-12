package implement_queue_using_stacks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyQueue(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		is := assert.New(t)
		obj := Constructor()
		obj.Push(1)
		obj.Push(2)
		is.Equal(1, obj.Peek())
		is.Equal(1, obj.Pop())
		is.False(obj.Empty())
	})

	t.Run("Example 2", func(t *testing.T) {
		is := assert.New(t)
		obj := Constructor()
		obj.Push(1)
		is.Equal(1, obj.Pop())
		is.True(obj.Empty())
	})

	t.Run("Example 3", func(t *testing.T) {
		queue := Constructor()
		queue.Push(1)
		queue.Push(2)
		assert.Equal(t, 1, queue.Peek())
		assert.Equal(t, 1, queue.Pop())
		assert.Equal(t, 2, queue.Pop())
		assert.True(t, queue.Empty())
	})
}
