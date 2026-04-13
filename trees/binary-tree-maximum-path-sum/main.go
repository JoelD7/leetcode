package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var dfs func(node *TreeNode) int
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	res := root.Val

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftSum := max(0, dfs(node.Left))
		rightSum := max(0, dfs(node.Right))

		res = max(res, leftSum+rightSum+node.Val)

		return max(leftSum, rightSum) + node.Val
	}

	dfs(root)

	return res
}
