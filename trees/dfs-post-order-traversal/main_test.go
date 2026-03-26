package main

import (
	"testing"

	"github.com/JoelD7/leetcode/trees/utils"
	"github.com/stretchr/testify/assert"
)

func TestPostOrderTraversalRecursive(t *testing.T) {
	t.Run("Complete tree", func(t *testing.T) {
		/*
		           26
		        /      \
		      19        33
		     /  \      /  \
		   12    24  31    40
		*/
		tree := utils.BuildTree([]*int{ptr(26), ptr(19), ptr(33), ptr(12), ptr(24), ptr(31), ptr(40)})
		assert.Equal(t, []int{12, 24, 19, 31, 40, 33, 26}, postOrderTraversalRecursive(tree))
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
		assert.Equal(t, []int{-12, -19, 40, 33, 26}, postOrderTraversalRecursive(tree))
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
		assert.Equal(t, []int{8, 12, 20, 19, 33, 26}, postOrderTraversalRecursive(tree))
	})
}

func TestPostOrderTraversalIterative(t *testing.T) {
	t.Run("Complete tree", func(t *testing.T) {
		/*
		           26
		        /      \
		      19        33
		     /  \      /  \
		   12    24  31    40
		*/
		tree := utils.BuildTree([]*int{ptr(26), ptr(19), ptr(33), ptr(12), ptr(24), ptr(31), ptr(40)})
		assert.Equal(t, []int{12, 24, 19, 31, 40, 33, 26}, postOrderTraversalIterative(tree))
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
		assert.Equal(t, []int{-12, -19, 40, 33, 26}, postOrderTraversalIterative(tree))
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
		assert.Equal(t, []int{8, 12, 20, 19, 33, 26}, postOrderTraversalIterative(tree))
	})
}

func ptr(i int) *int {
	return &i
}
