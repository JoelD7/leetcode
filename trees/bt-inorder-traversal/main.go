package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	return walk(root, []int{})
}

func walk(node *TreeNode, path []int) []int {
	if node == nil {
		return path
	}

	path = walk(node.Left, path)
	path = append(path, node.Val)
	path = walk(node.Right, path)

	return path
}
