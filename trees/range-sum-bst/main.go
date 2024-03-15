package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	path := walk(root, low, high, []int{})

	sum := 0
	for _, val := range path {
		sum += val
	}

	return sum
}

func walk(node *TreeNode, low, high int, path []int) []int {
	if node == nil {
		return path
	}

	if node.Val < low {
		path = walk(node.Right, low, high, path)
		return path
	}

	if node.Val > high {
		path = walk(node.Left, low, high, path)
		return path
	}

	path = append(path, node.Val)
	path = walk(node.Left, low, high, path)
	path = walk(node.Right, low, high, path)

	return path
}
