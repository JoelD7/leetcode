package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := 1
	getMax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var recurse func(node *TreeNode, level int) int

	recurse = func(node *TreeNode, level int) int {
		if node == nil {
			return getMax(level-1, max)
		}

		max = recurse(node.Left, level+1)
		max = recurse(node.Right, level+1)

		return max
	}

	max = recurse(root, 1)

	return max
}
