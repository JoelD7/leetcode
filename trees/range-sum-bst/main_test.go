package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRangeSumBST(t *testing.T) {
	c := require.New(t)

	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  15,
			Left: nil,
			Right: &TreeNode{
				Val:   18,
				Left:  nil,
				Right: nil,
			},
		},
	}

	c.Equal(32, rangeSumBST(root, 7, 15))
}
