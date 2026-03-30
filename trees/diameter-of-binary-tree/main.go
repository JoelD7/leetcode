package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	max := func(a, b int) int {
		if a > b {
			return a
		}

		return b
	}

	var getMaxDistance func(node *TreeNode) int
	getMaxDistance = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := getMaxDistance(node.Left)
		right := getMaxDistance(node.Right)

		res = max(res, left+right)

		return 1 + max(left, right)
	}

	getMaxDistance(root)

	return res
}
