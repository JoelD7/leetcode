package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var getDepth func(node *TreeNode) (int, bool)

	getDepth = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, true
		}

		left, balanced := getDepth(node.Left)
		if !balanced {
			return 0, false
		}

		right, balanced := getDepth(node.Right)
		if !balanced {
			return 0, false
		}

		if left-right > 1 || left-right < -1 {
			return 0, false
		}

		return max(left, right) + 1, true
	}

	_, balanced := getDepth(root)
	return balanced
}
