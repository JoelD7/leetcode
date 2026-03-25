package main

import (
	"testing"

	"github.com/JoelD7/leetcode/trees/utils"
	"github.com/stretchr/testify/assert"
)

func TestInOrderTraversalRecursive(t *testing.T) {
	t.Run("Complete tree", func(t *testing.T) {
		/*
		           26
		        /      \
		      19        33
		     /  \      /  \
		   12    24  31    40
		*/
		tree := utils.BuildTree([]*int{ptr(26), ptr(19), ptr(33), ptr(12), ptr(24), ptr(31), ptr(40)})
		assert.Equal(t, []int{12, 19, 24, 26, 31, 33, 40}, inOrderTraversalRecursive(tree))
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
		assert.Equal(t, []int{-12, -19, 26, 33, 40}, inOrderTraversalRecursive(tree))
	})

	t.Run("Very unbalanced tree", func(t *testing.T) {
		/*
		           26
		        /      \
		      19        33
		     /
		   12
		  /
		 8
		*/
		tree := utils.BuildTree([]*int{ptr(26), ptr(19), ptr(33), ptr(12), nil, nil, nil, ptr(8)})
		assert.Equal(t, []int{8, 12, 19, 26, 33}, inOrderTraversalRecursive(tree))
	})
}

func TestInOrderTraversalIterative(t *testing.T) {
	t.Run("Complete tree", func(t *testing.T) {
		/*
		           26
		        /      \
		      19        33
		     /  \      /  \
		   12    24  31    40
		*/
		tree := utils.BuildTree([]*int{ptr(26), ptr(19), ptr(33), ptr(12), ptr(24), ptr(31), ptr(40)})
		assert.Equal(t, []int{12, 19, 24, 26, 31, 33, 40}, inOrderTraversalIterative(tree))
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
		assert.Equal(t, []int{-12, -19, 26, 33, 40}, inOrderTraversalIterative(tree))
	})

	t.Run("Very unbalanced tree", func(t *testing.T) {
		/*
		           26
		        /      \
		      19        33
		     /
		   12
		  /
		 8
		*/
		tree := utils.BuildTree([]*int{ptr(26), ptr(19), ptr(33), ptr(12), nil, nil, nil, ptr(8)})
		assert.Equal(t, []int{8, 12, 19, 26, 33}, inOrderTraversalIterative(tree))
	})
}

func ptr(i int) *int { return &i }
