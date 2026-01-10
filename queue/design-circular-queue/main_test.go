package design_circular_queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		cQueue := Constructor(3)
		assert.True(t, cQueue.EnQueue(1))
		assert.True(t, cQueue.EnQueue(2))
		assert.True(t, cQueue.EnQueue(3))
		assert.False(t, cQueue.EnQueue(4))
		assert.Equal(t, 3, cQueue.Rear())
		assert.True(t, cQueue.IsFull())
		assert.True(t, cQueue.DeQueue())
		assert.True(t, cQueue.EnQueue(4))
		assert.Equal(t, 4, cQueue.Rear())
	})

	t.Run("Test case 2", func(t *testing.T) {
		cQueue := Constructor(3)
		assert.True(t, cQueue.EnQueue(2))
		assert.Equal(t, 2, cQueue.Rear())
		assert.Equal(t, 2, cQueue.Front())
		assert.True(t, cQueue.DeQueue())
		assert.Equal(t, -1, cQueue.Front())
		assert.False(t, cQueue.DeQueue())
		assert.Equal(t, -1, cQueue.Front())
		assert.True(t, cQueue.EnQueue(4))
		assert.True(t, cQueue.EnQueue(2))
		assert.True(t, cQueue.EnQueue(2))
		assert.False(t, cQueue.EnQueue(3))
	})
}
