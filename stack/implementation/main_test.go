package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := NewStack()

	assert.Equal(t, 0, s.Pop())

	s.Push(1) //1
	s.Push(2) //2 -> 1
	assert.Equal(t, 2, s.size)

	s.Push(3) //3 -> 2 -> 1
	assert.Equal(t, 3, s.size)
	assert.Equal(t, 3, s.Peek())
	assert.Equal(t, 3, s.size)

	assert.Equal(t, 3, s.Pop()) //2 -> 1
	assert.Equal(t, 2, s.size)

	s.Push(4) //4 -> 2 -> 1
	assert.Equal(t, 4, s.Peek())
	s.Push(5) //5 -> 4 -> 2 -> 1
	s.Push(6) //6 -> 5 -> 4 -> 2 -> 1
	assert.Equal(t, 6, s.Pop())
	assert.Equal(t, 5, s.Pop())
	assert.Equal(t, 4, s.Pop())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 1, s.Pop())
	assert.Equal(t, 0, s.size)
}
