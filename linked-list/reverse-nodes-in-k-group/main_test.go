package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseKGroup(t *testing.T) {
	t.Run("head = [1,2,3,4,5], k = 2", func(t *testing.T) {
		list := BuildList(1, 2, 3, 4, 5)

		assert.Equal(t, []int{2, 1, 4, 3, 5}, reverseKGroup(list, 2))
	})
}

func TestReverse(t *testing.T) {
	t.Run("list = [1,2,3,4]", func(t *testing.T) {
		list := BuildList(1, 2, 3, 4)
		revHead, revTail := reverse(list)

		assert.Equal(t, []int{4, 3, 2, 1}, listToArray(revHead))
		assert.NotNil(t, revTail)
		assert.Nil(t, revTail.Next)
		assert.Equal(t, 1, revTail.Val)
	})

	t.Run("list = [1,2,3]", func(t *testing.T) {
		list := BuildList(1, 2, 3)
		revHead, revTail := reverse(list)
		assert.Equal(t, []int{3, 2, 1}, listToArray(revHead))
		assert.NotNil(t, revTail)
		assert.Nil(t, revTail.Next)
		assert.Equal(t, 1, revTail.Val)
	})
}
