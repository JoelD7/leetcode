package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinChange(t *testing.T) {
	t.Run("basic example 1", func(t *testing.T) {
		coins := []int{1, 2, 5}
		amount := 11
		assert.Equal(t, 3, coinChange(coins, amount))
	})

	t.Run("no possible combination", func(t *testing.T) {
		coins := []int{2}
		amount := 3
		assert.Equal(t, -1, coinChange(coins, amount))
	})

	t.Run("amount is zero", func(t *testing.T) {
		coins := []int{1}
		amount := 0
		assert.Equal(t, 0, coinChange(coins, amount))
	})

	t.Run("single coin type", func(t *testing.T) {
		coins := []int{1}
		amount := 2
		assert.Equal(t, 2, coinChange(coins, amount))
	})

	t.Run("larger denominations", func(t *testing.T) {
		coins := []int{186, 419, 83, 408}
		amount := 6249
		assert.Equal(t, 20, coinChange(coins, amount))
		//6249/83=75 (remainder 24)
		//6249/186=33 (remainder 111)
		//6249/408=15 (remainder 129)
		//6249/419=14 (remainder 383)
	})
}
