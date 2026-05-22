package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthLargest(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		kthLargest := Constructor(3, []int{4, 5, 8, 2})
		assert.Equal(t, 4, kthLargest.Add(3))
		assert.Equal(t, 5, kthLargest.Add(5))
		assert.Equal(t, 5, kthLargest.Add(10))
		assert.Equal(t, 8, kthLargest.Add(9))
		assert.Equal(t, 8, kthLargest.Add(4))
	})

	t.Run("Test case 1", func(t *testing.T) {
		//["KthLargest","add","add","add","add","add"]
		//[[1,[]],[-3],[-2],[-4],[0],[4]

		kthLargest := Constructor(1, []int{})
		assert.Equal(t, -3, kthLargest.Add(-3))
		assert.Equal(t, -2, kthLargest.Add(-2))
		assert.Equal(t, -2, kthLargest.Add(-4))
		assert.Equal(t, 0, kthLargest.Add(0))
		assert.Equal(t, 4, kthLargest.Add(4))
	})
}
