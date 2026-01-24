package find_next_lower

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindNextLower(t *testing.T) {
	arr := make([]int, 100)

	for i := 0; i < 100; i++ {
		arr[i] = i + 1
	}

	next, iterations := findNextLower(100, arr)
	assert.Equal(t, 99, next)
	assert.Equal(t, 7, iterations)

	next, iterations = findNextLower(45, arr)
	assert.Equal(t, 44, next)
	assert.Equal(t, 7, iterations)

	next, iterations = findNextLower(2, arr)
	assert.Equal(t, 1, next)
	assert.Equal(t, 7, iterations)

	next, iterations = findNextLower(37, arr)
	assert.Equal(t, 36, next)
	assert.Equal(t, 7, iterations)

	next, iterations = findNextLower(66, arr)
	assert.Equal(t, 65, next)
	assert.Equal(t, 6, iterations)

	next, iterations = findNextLower(76, arr)
	assert.Equal(t, 75, next)
	assert.Equal(t, 6, iterations)
}
