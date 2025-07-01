package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	c := require.New(t)

	t.Run("7,1,5,3,6,4", func(t *testing.T) {
		prices := []int{7, 1, 5, 3, 6, 4}
		c.Equal(5, maxProfit(prices))
	})

}
