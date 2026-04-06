package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowestCommonAncestor(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		root := BuildTree([]*int{ptr(6), ptr(2), ptr(8), ptr(0), ptr(4), ptr(7), ptr(9), nil, nil, ptr(3), ptr(5)})
		p := &TreeNode{Val: 2}
		q := &TreeNode{Val: 8}

		assert.Equal(t, 6, lowestCommonAncestor(root, p, q).Val)
	})

	t.Run("Example 2", func(t *testing.T) {
		root := BuildTree([]*int{ptr(6), ptr(2), ptr(8), ptr(0), ptr(4), ptr(7), ptr(9), nil, nil, ptr(3), ptr(5)})
		p := &TreeNode{Val: 2}
		q := &TreeNode{Val: 4}

		assert.Equal(t, 2, lowestCommonAncestor(root, p, q).Val)
	})

	t.Run("Test case 4", func(t *testing.T) {
		//[3,1,4,null,2]
		root := BuildTree([]*int{ptr(3), ptr(1), ptr(4), nil, ptr(2)})
		p := &TreeNode{Val: 2}
		q := &TreeNode{Val: 4}

		assert.Equal(t, 3, lowestCommonAncestor(root, p, q).Val)
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

func ptr(i int) *int { return &i }
