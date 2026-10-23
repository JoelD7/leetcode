package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubtree(t *testing.T) {
	t.Run("Standard match in left subtree", func(t *testing.T) {
		// root = [3,4,5,1,2]
		root := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			Right: &TreeNode{Val: 5},
		}

		// subRoot = [4,1,2]
		subRoot := &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		}

		assert.True(t, isSubtree(root, subRoot))
	})

	t.Run("Match interrupted by extra leaf node", func(t *testing.T) {
		// root = [3,4,5,1,2,null,null,null,null,0]
		root := &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:  4,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 0},
				},
			},
			Right: &TreeNode{Val: 5},
		}

		// subRoot = [4,1,2]
		subRoot := &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		}

		assert.False(t, isSubtree(root, subRoot))
	})

	t.Run("Subroot is exactly the same as root", func(t *testing.T) {
		root := &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 3},
		}

		subRoot := &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 3},
		}

		assert.True(t, isSubtree(root, subRoot))
	})

	t.Run("Single node match", func(t *testing.T) {
		root := &TreeNode{Val: 10}
		subRoot := &TreeNode{Val: 10}

		assert.True(t, isSubtree(root, subRoot))
	})

	t.Run("Single node mismatch", func(t *testing.T) {
		root := &TreeNode{Val: 10}
		subRoot := &TreeNode{Val: 15}

		assert.False(t, isSubtree(root, subRoot))
	})

	t.Run("Similar values but wrong structure", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 1},
			},
		}

		// subRoot has right child instead of left
		subRoot := &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 1},
		}

		assert.False(t, isSubtree(root, subRoot))
	})

	t.Run("root =[10,5,null,4], subRoot =[10,4,null,null,5]", func(t *testing.T) {
		// root = [10,5,null,4]
		root := &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
				},
			},
		}

		// subRoot = [10,4,null,null,5]
		subRoot := &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 5,
				},
			},
		}

		assert.False(t, isSubtree(root, subRoot))
	})
}
