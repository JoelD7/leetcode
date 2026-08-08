package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCostClimbingStairs(t *testing.T) {
	// Standard LeetCode Example 1
	assert.Equal(t, 15, minCostClimbingStairs([]int{10, 15, 20}), "Example 1 should return 15")

	// Standard LeetCode Example 2
	assert.Equal(t, 6, minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}), "Example 2 should return 6")

	// Minimum constraint length (2 elements) - starting from index 0 is cheaper
	assert.Equal(t, 10, minCostClimbingStairs([]int{10, 20}), "Two elements - index 0 cheaper")

	// Minimum constraint length (2 elements) - starting from index 1 is cheaper
	assert.Equal(t, 10, minCostClimbingStairs([]int{20, 10}), "Two elements - index 1 cheaper")

	// Array with zero costs
	assert.Equal(t, 0, minCostClimbingStairs([]int{0, 0, 0, 0}), "Zero costs should return 0")

	// Strictly increasing costs
	assert.Equal(t, 12, minCostClimbingStairs([]int{2, 4, 6, 8, 10}), "Strictly increasing costs")

	// Strictly decreasing costs
	assert.Equal(t, 12, minCostClimbingStairs([]int{10, 8, 6, 4, 2}), "Strictly decreasing costs")
}
