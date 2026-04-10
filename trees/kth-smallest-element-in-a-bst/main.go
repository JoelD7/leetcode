package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	arr := make([]int, 0)
	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}

		smallest := dfs(node.Left)
		if smallest != -1 {
			return smallest
		}

		arr = append(arr, node.Val)
		if len(arr) == k {
			return arr[len(arr)-1]
		}

		return dfs(node.Right)
	}

	return dfs(root)
}
