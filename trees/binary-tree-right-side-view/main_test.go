package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRightSideView(t *testing.T) {
	t.Run("Standard tree with both sides", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 5},
			},
			Right: &TreeNode{
				Val:   3,
				Right: &TreeNode{Val: 4},
			},
		}
		assert.Equal(t, []int{1, 3, 4}, rightSideView(root))
	})

	t.Run("Right child only", func(t *testing.T) {
		root := &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 3},
		}
		assert.Equal(t, []int{1, 3}, rightSideView(root))
	})

	t.Run("Empty tree", func(t *testing.T) {
		var root *TreeNode = nil
		assert.Empty(t, rightSideView(root))
	})

	t.Run("Left side is deeper than right side", func(t *testing.T) {
		// Tree:
		//     1
		//    / \
		//   2   3
		//  /
		// 4
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 4},
			},
			Right: &TreeNode{Val: 3},
		}
		assert.Equal(t, []int{1, 3, 4}, rightSideView(root))
	})

	t.Run("Single node tree", func(t *testing.T) {
		root := &TreeNode{Val: 10}
		assert.Equal(t, []int{10}, rightSideView(root))
	})
}
