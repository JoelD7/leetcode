package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	t.Run("Test 1: General unsorted array", func(t *testing.T) {
		items := []int{5, 2, 8, 0, 1, 6, 3, 9, 7, 4}
		assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, mergeSort(items))
	})

	t.Run("Test 2: Empty array", func(t *testing.T) {
		items := []int{}
		assert.Equal(t, []int{}, mergeSort(items))
	})

	t.Run("Test 3: Single element array", func(t *testing.T) {
		items := []int{42}
		assert.Equal(t, []int{42}, mergeSort(items))
	})

	t.Run("Test 4: Already sorted array", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5, 6}
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, mergeSort(items))
	})

	t.Run("Test 5: Reverse sorted array", func(t *testing.T) {
		items := []int{9, 8, 7, 6, 5}
		assert.Equal(t, []int{5, 6, 7, 8, 9}, mergeSort(items))
	})

	t.Run("Test 6: Array with negative numbers and duplicates", func(t *testing.T) {
		items := []int{3, -1, 4, 3, -1, 0}
		assert.Equal(t, []int{-1, -1, 0, 3, 3, 4}, mergeSort(items))
	})
}

func TestQuicksort(t *testing.T) {
	t.Run("Test 1: General unsorted array", func(t *testing.T) {
		items := []int{5, 2, 8, 0, 1, 6, 3, 9, 7, 4}
		assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, quickSort(items, 0, len(items)-1))
	})

	t.Run("Test 2: Empty array", func(t *testing.T) {
		items := []int{}
		assert.Equal(t, []int{}, quickSort(items, 0, len(items)-1))
	})

	t.Run("Test 3: Single element array", func(t *testing.T) {
		items := []int{42}
		assert.Equal(t, []int{42}, quickSort(items, 0, len(items)-1))
	})

	t.Run("Test 4: Already sorted array", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5, 6}
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, quickSort(items, 0, len(items)-1))
	})

	t.Run("Test 5: Reverse sorted array", func(t *testing.T) {
		items := []int{9, 8, 7, 6, 5}
		assert.Equal(t, []int{5, 6, 7, 8, 9}, quickSort(items, 0, len(items)-1))
	})

	t.Run("Test 6: Array with negative numbers and duplicates", func(t *testing.T) {
		items := []int{3, -1, 4, 3, -1, 0}
		assert.Equal(t, []int{-1, -1, 0, 3, 3, 4}, quickSort(items, 0, len(items)-1))
	})
}
