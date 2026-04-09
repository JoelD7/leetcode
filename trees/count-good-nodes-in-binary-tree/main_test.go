package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoodNodes(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		tree := buildTree([]*int{ptr(3), ptr(1), ptr(4), ptr(3), nil, ptr(1), ptr(5)})
		assert.Equal(t, 4, goodNodes(tree))
	})

	t.Run("Example 2", func(t *testing.T) {
		tree := buildTree([]*int{ptr(3), ptr(3), nil, ptr(4), ptr(2)})
		assert.Equal(t, 3, goodNodes(tree))
	})

	t.Run("Example 3", func(t *testing.T) {
		tree := buildTree([]*int{ptr(1)})
		assert.Equal(t, 1, goodNodes(tree))
	})
}

func buildTree(values []*int) *TreeNode {
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
