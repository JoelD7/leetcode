package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCarFleet(t *testing.T) {
	c := require.New(t)

	c.Equal(3, carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))
	c.Equal(1, carFleet(10, []int{3}, []int{3}))
	c.Equal(1, carFleet(100, []int{0, 2, 4}, []int{4, 2, 1}))
}
