package main

import (
	"testing"

	"github.com/JoelD7/leetcode/trees/utils"
	"github.com/stretchr/testify/assert"
)

func TestPreOrderTraversalRecursive(t *testing.T) {
	t.Run("Complete tree", func(t *testing.T) {
		/*
		           26
		        /      \
		      19        33
		     /  \      /  \
		   12    24  31    40
		*/
		tree := utils.BuildTree([]*int{ptr(26), ptr(19), ptr(33), ptr(12), ptr(24), ptr(31), ptr(40)})
		assert.Equal(t, []int{26, 19, 12, 24, 33, 31, 40}, preOrderTraversalRecursive(tree))
	})

	t.Run("Non-complete tree", func(t *testing.T) {
		/*
		           26
		        /      \
		     -19        33
		     /           \
		  -12             40
		*/
		tree := utils.BuildTree([]*int{ptr(26), ptr(-19), ptr(33), ptr(-12), nil, nil, ptr(40)})
		assert.Equal(t, []int{26, -19, -12, 33, 40}, preOrderTraversalRecursive(tree))
	})

	t.Run("Very unbalanced tree", func(t *testing.T) {
		/*
		           26
		        /      \
		      19        33
		     /  \
		   12   20
		  /
		 8
		*/
		tree := utils.BuildTree([]*int{ptr(26), ptr(19), ptr(33), ptr(12), ptr(20), nil, nil, ptr(8)})
		assert.Equal(t, []int{26, 19, 12, 8, 20, 33}, preOrderTraversalRecursive(tree))
	})
}

func TestPreOrderTraversalIterative(t *testing.T) {
	t.Run("Complete tree", func(t *testing.T) {
		/*
		           26
		        /      \
		      19        33
		     /  \      /  \
		   12    24  31    40
		*/
		tree := utils.BuildTree([]*int{ptr(26), ptr(19), ptr(33), ptr(12), ptr(24), ptr(31), ptr(40)})
		assert.Equal(t, []int{26, 19, 12, 24, 33, 31, 40}, preOrderTraversalIterative(tree))
	})

	t.Run("Non-complete tree", func(t *testing.T) {
		/*
		           26
		        /      \
		     -19        33
		     /           \
		  -12             40
		*/
		tree := utils.BuildTree([]*int{ptr(26), ptr(-19), ptr(33), ptr(-12), nil, nil, ptr(40)})
		assert.Equal(t, []int{26, -19, -12, 33, 40}, preOrderTraversalIterative(tree))
	})

	t.Run("Very unbalanced tree", func(t *testing.T) {
		/*
		           26
		        /      \
		      19        33
		     /  \
		   12   20
		  /
		 8
		*/
		tree := utils.BuildTree([]*int{ptr(26), ptr(19), ptr(33), ptr(12), ptr(20), nil, nil, ptr(8)})
		assert.Equal(t, []int{26, 19, 12, 8, 20, 33}, preOrderTraversalIterative(tree))
	})
}

func ptr(i int) *int {
	return &i
}
