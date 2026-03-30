package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const nilValue = -10001

// This is from when I first solved this problem on march 2024. It's not optimal but I left here for historic reasons.
func isSameTreeDeprecated(p *TreeNode, q *TreeNode) bool {
	pPath := walk(p, []int{})
	qPath := walk(q, []int{})

	return compareSlices(pPath, qPath)
}

func walk(node *TreeNode, path []int) []int {
	if node == nil {
		path = append(path, nilValue)
		return path
	}

	path = append(path, node.Val)
	path = walk(node.Left, path)
	path = walk(node.Right, path)

	return path
}

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
