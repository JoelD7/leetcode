package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var arr []int
	var inOrderTraversal func(node *TreeNode) bool

	inOrderTraversal = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		if !inOrderTraversal(node.Left) {
			return false
		}

		if len(arr) > 0 && node.Val <= (arr)[len(arr)-1] {
			return false
		}

		arr = append(arr, node.Val)

		if !inOrderTraversal(node.Right) {
			return false
		}

		return true
	}

	return inOrderTraversal(root)
}
