package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMediaFromDataStream(t *testing.T) {
	t.Run("Test case 1", func(t *testing.T) {
		mf := Constructor()
		mf.AddNum(-1)
		assert.Equal(t, -1.0, mf.FindMedian())
		mf.AddNum(-2)
		assert.Equal(t, -1.5, mf.FindMedian())
		mf.AddNum(-3)
		assert.Equal(t, -2.0, mf.FindMedian())
		mf.AddNum(-4)
		assert.Equal(t, -2.5, mf.FindMedian())
		mf.AddNum(-5)
		assert.Equal(t, -3.0, mf.FindMedian())
	})
}
