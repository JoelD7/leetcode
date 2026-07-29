package queue_impl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := NewQueue()

	assert.True(t, q.IsEmpty())
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	assert.Equal(t, 4, q.size)
	assert.False(t, q.IsEmpty())

	val, err := q.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 1, val)
	assert.Equal(t, 4, q.size)

	val, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 1, val)
	assert.Equal(t, 3, q.size)

	val, err = q.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 2, val)
	assert.Equal(t, 3, q.size)

	q.Enqueue(5)
	q.Enqueue(6)
	q.Enqueue(7)
	assert.Equal(t, 6, q.size)

	val, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 2, val)
	assert.Equal(t, 5, q.size)

	val, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 3, val)
	assert.Equal(t, 4, q.size)

	val, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 4, val)
	assert.Equal(t, 3, q.size)

	val, err = q.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 5, val)

	val, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 5, val)

	val, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 6, val)

	val, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 7, val)

	val, err = q.Dequeue()
	assert.Error(t, err)
}
