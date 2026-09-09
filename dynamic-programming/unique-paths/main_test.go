package unique_paths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	t.Run("Leetcode Example 1: m = 3, n = 7", func(t *testing.T) {
		assert.Equal(t, 28, uniquePaths(3, 7))
	})

	t.Run("Leetcode Example 2: m = 3, n = 2", func(t *testing.T) {
		assert.Equal(t, 3, uniquePaths(3, 2))
	})

	t.Run("Smallest possible grid: m = 1, n = 1", func(t *testing.T) {
		assert.Equal(t, 1, uniquePaths(1, 1))
	})

	t.Run("Single row grid: m = 1, n = 10", func(t *testing.T) {
		assert.Equal(t, 1, uniquePaths(1, 10))
	})

	t.Run("Single column grid: m = 10, n = 1", func(t *testing.T) {
		assert.Equal(t, 1, uniquePaths(10, 1))
	})

	t.Run("Square grid: m = 3, n = 3", func(t *testing.T) {
		assert.Equal(t, 6, uniquePaths(3, 3))
	})

	t.Run("Larger square grid: m = 4, n = 4", func(t *testing.T) {
		assert.Equal(t, 20, uniquePaths(4, 4))
	})

	t.Run("Inverted dimensions: m = 7, n = 3", func(t *testing.T) {
		assert.Equal(t, 28, uniquePaths(7, 3))
	})

	t.Run("m = 23, n = 12", func(t *testing.T) {
		assert.Equal(t, 193536720, uniquePaths(23, 12))
	})
}

//Invalid paths
//6 0
//6 1
//6 1
//5 2
//6 1
//5 2
//5 2
//4 2
//6 1
//5 2
//5 2
//4 2
//5 2
//4 2
//3 2
//6 1
//5 2
//5 2
//4 2
//5 2
//4 2
//3 2
//5 2
//4 2
//3 2
//2 2
//6 1
//5 2
//5 2
//4 2
//5 2
//4 2
//3 2
//5 2
//4 2
//3 2
//2 2
//5 2
//4 2
//3 2
//2 2
//1 2
//6 1
//5 2
//5 2
//4 2
//5 2
//4 2
//3 2
//5 2
//4 2
//3 2
//2 2
//5 2
//4 2
//3 2
//2 2
//1 2
//5 2
//4 2
//3 2
//2 2
//1 2
//0 2
