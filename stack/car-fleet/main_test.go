package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCarFleet(t *testing.T) {
	c := require.New(t)

	t.Run("target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]", func(t *testing.T) {
		c.Equal(3, carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))
	})

	t.Run("target = 10, position = [3], speed = [3]", func(t *testing.T) {
		c.Equal(1, carFleet(10, []int{3}, []int{3}))
	})

	t.Run("target = 100, position = [0,2,4], speed = [4,2,1]", func(t *testing.T) {
		c.Equal(1, carFleet(100, []int{0, 2, 4}, []int{4, 2, 1}))
	})

	t.Run("target = 10, position = [6,8], speed = [3.2]", func(t *testing.T) {
		c.Equal(2, carFleet(10, []int{6, 8}, []int{3, 2}))
	})

}
