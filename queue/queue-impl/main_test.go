package queue_impl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()

	assert.True(t, q.isEmpty())
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	assert.Equal(t, 4, q.size)
	assert.False(t, q.isEmpty())
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
