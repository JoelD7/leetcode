package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidBST(t *testing.T) {
	t.Run("Valid simple tree [2,1,3]", func(t *testing.T) {
		root := &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		}
		assert.Equal(t, true, isValidBST(root))
	})

	t.Run("Invalid tree right child less than root [5,1,4,null,null,3,6]", func(t *testing.T) {
		root := &TreeNode{
			Val:  5,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 6},
			},
		}
		assert.Equal(t, false, isValidBST(root))
	})

	t.Run("Single node [1]", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		assert.Equal(t, true, isValidBST(root))
	})

	t.Run("Empty tree []", func(t *testing.T) {
		assert.Equal(t, true, isValidBST(nil))
	})

	t.Run("Invalid tree deep subtree violation [5,4,6,null,null,3,7]", func(t *testing.T) {
		root := &TreeNode{
			Val:  5,
			Left: &TreeNode{Val: 4},
			Right: &TreeNode{
				Val:   6,
				Left:  &TreeNode{Val: 3}, // 3 is less than the root 5, which violates BST rules
				Right: &TreeNode{Val: 7},
			},
		}
		assert.Equal(t, false, isValidBST(root))
	})

	t.Run("Valid negative values [-2,-3,-1]", func(t *testing.T) {
		root := &TreeNode{
			Val:   -2,
			Left:  &TreeNode{Val: -3},
			Right: &TreeNode{Val: -1},
		}
		assert.Equal(t, true, isValidBST(root))
	})
}
