package main

import (
	"testing"

	"github.com/JoelD7/leetcode/trees/utils"
	"github.com/stretchr/testify/assert"
)

func TestBFSTraversalRecursive(t *testing.T) {
	t.Run("Complete tree", func(t *testing.T) {
		/*
		           26
		        /      \
		      19        33
		     /  \      /  \
		   12    24  31    40
		*/
		tree := utils.BuildTree([]*int{ptr(26), ptr(19), ptr(33), ptr(12), ptr(24), ptr(31), ptr(40)})
		assert.Equal(t, []int{26, 19, 33, 12, 24, 31, 40}, flatten(bfsTraversalRecursive(tree)))
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
		assert.Equal(t, []int{26, -19, 33, -12, 40}, flatten(bfsTraversalRecursive(tree)))
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
		assert.Equal(t, []int{26, 19, 33, 12, 20, 8}, flatten(bfsTraversalRecursive(tree)))
	})
}

func TestBFSTraversalIterative(t *testing.T) {
	t.Run("Complete tree", func(t *testing.T) {
		/*
		           26
		        /      \
		      19        33
		     /  \      /  \
		   12    24  31    40
		*/
		tree := utils.BuildTree([]*int{ptr(26), ptr(19), ptr(33), ptr(12), ptr(24), ptr(31), ptr(40)})
		assert.Equal(t, []int{26, 19, 33, 12, 24, 31, 40}, flatten(bfsTraversalIterative(tree)))
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
		assert.Equal(t, []int{26, -19, 33, -12, 40}, flatten(bfsTraversalIterative(tree)))
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
		assert.Equal(t, []int{26, 19, 33, 12, 20, 8}, flatten(bfsTraversalIterative(tree)))
	})
}

func flatten(arr [][]int) []int {
	var result []int

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			result = append(result, arr[i][j])
		}
	}

	return result
}

func ptr(i int) *int {
	return &i
}
