package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	count := 0

	var dfs func(node *TreeNode, largest *TreeNode)
	dfs = func(node *TreeNode, largest *TreeNode) {
		if node == nil {
			return
		}

		if node.Val >= largest.Val {
			count++
			dfs(node.Left, node)
			dfs(node.Right, node)
			return
		}

		dfs(node.Left, largest)
		dfs(node.Right, largest)
	}

	dfs(root, root)

	return count
}
