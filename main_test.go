package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrefixProduct(t *testing.T) {
	c := require.New(t)

	t.Run("[1,2,3,4]", func(t *testing.T) {
		c.EqualValues([]int{1, 2, 6, 24}, prefixProduct([]int{1, 2, 3, 4}))
	})

	t.Run("[4,2,5,3]", func(t *testing.T) {
		c.EqualValues([]int{4, 8, 40, 120}, prefixProduct([]int{4, 2, 5, 3}))
	})
}

func TestSufixProduct(t *testing.T) {
	c := require.New(t)

	t.Run("[1,2,3,4]", func(t *testing.T) {
		c.EqualValues([]int{24, 24, 12, 4}, sufixProduct([]int{1, 2, 3, 4}))
	})
}
