package top_k_frequent_elements

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	c := require.New(t)

	t.Run("nums = [1,1,1,2,2,3], k = 2", func(t *testing.T) {
		c.Equal([]int{1, 2}, topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	})

	t.Run("nums = [1], k = 1", func(t *testing.T) {
		c.Equal([]int{1}, topKFrequent([]int{1}, 1))
	})

	t.Run("nums = [1,2,1,2,1,2,3,1,3,2], k = 2", func(t *testing.T) {
		c.Equal([]int{1, 2}, topKFrequent([]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2))
	})

}
