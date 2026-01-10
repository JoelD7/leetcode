package design_circular_queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		cQueue := Constructor(3)
		assert.True(t, cQueue.EnQueue(1))  // return True
		assert.True(t, cQueue.EnQueue(2))  // return True
		assert.True(t, cQueue.EnQueue(3))  // return True
		assert.False(t, cQueue.EnQueue(4)) // return False
		assert.Equal(t, 3, cQueue.Rear())  // return 3
		assert.True(t, cQueue.IsFull())    // return True
		assert.True(t, cQueue.DeQueue())   // return True
		assert.True(t, cQueue.EnQueue(4))  // return True
		assert.Equal(t, 4, cQueue.Rear())  // return 4
	})

	//["MyCircularQueue","enQueue","Rear","Front","deQueue","Front","deQueue"."Front", "enQueue","enQueue","enQueue","enQueue"]
	//[[3],[2],[],[],[],[],[],[],[4],[2],[2],[3]]
	//[null,true,2,2,true ,-1,false,-1, true,true,true,false]
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
