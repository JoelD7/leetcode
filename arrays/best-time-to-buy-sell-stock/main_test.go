package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxProfit(t *testing.T) {
	c := require.New(t)

	t.Run("7,1,5,3,6,4", func(t *testing.T) {
		prices := []int{7, 1, 5, 3, 6, 4}
		c.Equal(5, maxProfit(prices))
	})

	t.Run("2,1,4", func(t *testing.T) {
		c.Equal(3, maxProfit([]int{2, 1, 4}))
	})

	t.Run("2,1,4,5,2,9,7", func(t *testing.T) {
		c.Equal(8, maxProfit([]int{2, 1, 4, 5, 2, 9, 7}))
	})
}
