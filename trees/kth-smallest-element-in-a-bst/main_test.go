package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthSmallest(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		root := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
			Right: &TreeNode{Val: 4},
		}
		assert.Equal(t, 1, kthSmallest(root, 1))
	})

	t.Run("example 2", func(t *testing.T) {
		root := &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 1},
				},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{Val: 6},
		}
		assert.Equal(t, 3, kthSmallest(root, 3))
	})

	t.Run("single node tree", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		assert.Equal(t, 1, kthSmallest(root, 1))
	})
}
