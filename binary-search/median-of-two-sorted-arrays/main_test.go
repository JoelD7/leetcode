package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	t.Run("nums1 = [1, 3], nums2 = [2]", func(t *testing.T) {
		assert.Equal(t, 2.0, findMedianSortedArrays([]int{1, 3}, []int{2}))
	})

	t.Run("nums1 = [1, 2], nums2 = [3, 4]", func(t *testing.T) {
		assert.Equal(t, 2.5, findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	})

	t.Run("empty nums1, nums2 has one element", func(t *testing.T) {
		assert.Equal(t, 1.0, findMedianSortedArrays([]int{}, []int{1}))
	})

	t.Run("nums1 has one element, empty nums2", func(t *testing.T) {
		assert.Equal(t, 2.0, findMedianSortedArrays([]int{2}, []int{}))
	})

	t.Run("arrays with duplicate elements", func(t *testing.T) {
		assert.Equal(t, 1.0, findMedianSortedArrays([]int{1, 1}, []int{1, 2, 3}))
	})

	t.Run("arrays with negative numbers", func(t *testing.T) {
		assert.Equal(t, 3.0, findMedianSortedArrays([]int{-5, 3, 6, 12, 15}, []int{-12, -10, -6, -3, 4, 10}))
	})

	t.Run("arrays with all same elements", func(t *testing.T) {
		assert.Equal(t, 2.0, findMedianSortedArrays([]int{2, 2, 2}, []int{2, 2, 2, 2}))
	})

	t.Run("completely disjoint arrays", func(t *testing.T) {
		assert.Equal(t, 3.5, findMedianSortedArrays([]int{1, 2, 3}, []int{4, 5, 6}))
	})

	t.Run("emtpy second array", func(t *testing.T) {
		assert.Equal(t, 6.0, findMedianSortedArrays([]int{4, 5, 6, 8, 9}, []int{}))
	})
}
