package main

import (
	"testing"

	"github.com/JoelD7/leetcode/trees/utils"
	"github.com/stretchr/testify/assert"
)

var ptr = utils.Ptr

func TestIsValid(t *testing.T) {
	t.Run("Test case 1", func(t *testing.T) {
		tree := BuildTree([]*int{ptr(5), ptr(4), ptr(6), nil, nil, ptr(3), ptr(7)})
		assert.False(t, isValidBST(tree))
	})
}

func BuildTree(values []*int) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *values[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		current := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(values) && values[i] != nil {
			current.Left = &TreeNode{Val: *values[i]}
			queue = append(queue, current.Left)
		}
		i++

		// Right child
		if i < len(values) && values[i] != nil {
			current.Right = &TreeNode{Val: *values[i]}
			queue = append(queue, current.Right)
		}
		i++
	}
	return root
}
