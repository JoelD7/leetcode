package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var isSame func(root *TreeNode, subRoot *TreeNode) bool
	isSame = func(root *TreeNode, subRoot *TreeNode) bool {
		if root == nil && subRoot == nil {
			return true
		}

		if root == nil || subRoot == nil {
			return false
		}

		if root.Val != subRoot.Val {
			return false
		}

		return isSame(root.Left, subRoot.Left) && isSame(root.Right, subRoot.Right)
	}

	if root == nil {
		return false
	}

	if isSame(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
