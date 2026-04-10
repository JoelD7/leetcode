package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var count int
	smallest := -1
	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		//I think this will never happen. Check later
		if node == nil {
			return
		}

		if node.Left != nil && node.Left.Val < node.Val {
			dfs(node.Left)
		}

		if count < k && node.Right != nil {
			dfs(node.Right)
		}

		count++
		if count == k && smallest == -1 {
			smallest = node.Val
		}
	}

	dfs(root)

	return smallest
}
